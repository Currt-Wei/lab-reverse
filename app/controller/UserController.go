package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"lab-reverse/app/common"
	"lab-reverse/app/middleware"
	"lab-reverse/app/model"
	"lab-reverse/app/service"
	"lab-reverse/constant"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type LoginResult struct {
	Name  string
	Token string
}

// 定义一个普通controller函数，作为一个验证接口逻辑
func TestToken(c *gin.Context) {
	// 上面我们在JWTAuth()中间中将'claims'写入到gin.Context的指针对象中，因此在这里可以将之解析出来
	claims := c.MustGet("claims").(*middleware.CustomClaims)

	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
		return
	}



	c.JSON(http.StatusOK, gin.H{
		"status": -1,
		"msg":    "token失效",
		"data":   nil,
	})

	return
}

// Register 用户注册
// @Summary 学生/教师注册
// @Tags Register
// @Accept json
// @Produce json
// @Param	user  body	model.User  true  "Add user"
// @Success 200 {string} string
// @Router /register [post]
func Register(ctx *gin.Context) {
	var db = model.DB
	var tmp = struct {
		model.User
		Code string
	}{}
	// 绑定数据
	if err := ctx.Bind(&tmp); err != nil {
		log.Println("[register]注册绑定数据出错")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"data": nil,
			"msg": "注册绑定数据出错",
		})
		return
	}
	stu := tmp.User

	// 验证邮箱验证码是否正确
	emailCode, found := constant.EmailCache.Get("user_" + stu.Email)
	if !found || tmp.Code != emailCode {
		log.Println("[register]验证码错误")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": constant.CodeError,
			"data": nil,
			"msg": "验证码错误",
		})
		return
	}

	// 验证字段正确性
	validate := validator.New()
	err := validate.Struct(&stu)
	if err != nil {
		log.Println("[register]字段验证错误")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"data": nil,
			"msg": err.Error(),
		})
		return
	}

	// 判断学号是否重复
	var temp model.User
	db.Where("account", stu.Account).Find(&temp)
	if temp.Id != 0 {
		log.Println("[register]该用户已注册")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"data": nil,
			"msg": "该用户已注册",
		})
		return
	}

	// 密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(stu.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[register]密码加密失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"data": nil,
			"msg": "系统出错",
		})
		return
	}
	stu.Password = string(password)
	err = db.Create(&stu).Error
	if err != nil {
		log.Println("[register]创建用户失败, err:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"data": nil,
			"msg": "创建用户失败",
		})
		return
	}
	// 删掉邮箱验证码
	constant.EmailCache.Delete("stu_" + stu.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": nil,
		"msg": "注册成功",
	})
}

func Login(ctx *gin.Context) {
	var db = model.DB
	var loginUser model.User
	// 获取账号密码
	if err := ctx.Bind(&loginUser); err != nil {
		log.Println("[login]绑定数据出错")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"data": nil,
			"msg": "绑定数据出错",
		})
		return
	}
	// 查出当前用户
	var temp model.User
	db.Where("account", loginUser.Account).Find(&temp)
	if temp.Id == 0 {
		log.Println("[login]该用户不存在")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"data": nil,
			"msg": "账号或密码错误",
		})
		return
	}
	// 对比密码是否错误
	err := bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(loginUser.Password))
	if err != nil {
		log.Println("[login]密码错误")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"data": nil,
			"msg": "账号或密码错误",
		})
		return
	}
	// 查看用户是否是启用状态
	if temp.Enable == 0 {
		log.Println("[login]该用户是禁用状态")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": constant.DisableUserError,
			"data": nil,
			"msg": "该用户是禁用状态，请联系管理员",
		})
		return
	}
	// 登录成功，返回token
	token, err := common.ReleaseToken(temp)
	if err != nil {
		log.Println("[login]生成token出错")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"data": nil,
			"msg": "生成token出错",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
			"role": temp.RoleId,
		},
		"msg": "登录成功",
	})
}

// Login 登陆
//func Login(c *gin.Context) {
//	var u model.User
//	if err := c.ShouldBindJSON(&u); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status": constant.LoginFail,
//			"msg":    "登录失败",
//			"data":   err.Error(),
//		})
//		return
//	}
//
//	//TODO 查找数据库
//	user, err := service.FindUserByEmail(u.Email)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"status": constant.LoginFail,
//			"msg":    err.Error(),
//			"data":   "登录失败",
//		})
//		return
//	}
//
//	// 密码错误
//	if u.Password != user.Password {
//		c.JSON(http.StatusOK, gin.H{
//			"status": constant.LoginFail,
//			"msg":    "登录失败",
//			"data":   err.Error(),
//		})
//		return
//	}
//
//	token := generateToken(c, *user)
//	var role string
//	// 用户角色
//	//if len(user.Role) == 0 {
//	//	role = "user"
//	//} else {
//	//	role = user.Role[0].RoleName
//	//}
//	c.JSON(http.StatusOK, gin.H{
//		"status": constant.LoginSuccess,
//		"msg":    "登陆成功",
//		"data": gin.H{
//			"token": token,
//			"role":  role,
//		},
//	})
//
//	return
//}

// token生成器
func generateToken(c *gin.Context, user model.User) string {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := middleware.NewJWT()
	// 构造用户claims信息(负荷)
	claims := middleware.CustomClaims{
		user.Email,
		user.Password,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			// 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 36000),
			// 签名过期时间
			Issuer: "lab-reverse",
			// 签名颁发者
		},
	}
	// 根据claims生成token对象
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.LoginFail,
			"msg":    err.Error(),
			"data":   nil,
		})
	}

	return token
}

func GetInfoByAccount(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.GetInfoByEmailFail,
			"msg":    "查询失败",
			"data":   err.Error(),
		})
		return
	}

	//TODO 查找数据库
	user, err := service.FindUserByAccount(u.Account)

	fmt.Println("GetInfoByEmail 请求成功")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.GetInfoByEmailFail,
			"msg":    err.Error(),
			"data":   "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.GetInfoByEmailSuccess,
		"msg":    "查询成功",
		"data": 	user,
	})

	return
}

func GetReserveInfo(c *gin.Context){



	// 获取分页数据
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	account := c.Query("account")

	//users, err:=service.FindAllUser(data,limit,page)
	offset := (page - 1) * limit

	db:=model.DB
	db = db.Limit(limit).Offset(offset)

	var reserveInfo []model.Reservation
	err := db.Where("account = ?", account).Find(&reserveInfo).Error

	var todo []model.Reservation
	var done []model.Reservation
 	for _,info :=range reserveInfo {
		interval :=strings.Split(info.TimeInterval,"-")
		t,_ := time.ParseInLocation("2006-01-02 15:04:05", info.ReserveDate+" "+interval[0],time.Local)
		if t.After(time.Now()){
			todo = append(todo,info)
		} else {
			done = append(done, info)
		}


	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.GetReserveInfoFail,
			"msg":    err.Error(),
			"data":   "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.GetReserveInfoSuccess,
		"msg":    "查询成功",
		//"data": 	reserveInfo,
		"todo": 	todo,
		"done":		done,
	})

	return
}

