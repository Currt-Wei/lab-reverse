package controller

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "golang.org/x/crypto/bcrypt"
	"lab-reverse/app/common"
	"lab-reverse/app/dto"
	"lab-reverse/app/middleware"
	"lab-reverse/app/model"
	"lab-reverse/app/model/response"
	"lab-reverse/app/service"
	"lab-reverse/constant"
	"lab-reverse/global"
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

// SendVerificationCode 发送邮箱验证码
func SendVerificationCode(ctx *gin.Context) {
	// 获取邮箱
	b, err := ctx.GetRawData()  // 从c.Request.Body读取请求数据
	if err != nil {
		log.Println("[register]获取请求体失败")
		response.FailWithDetailed("500", nil, "获取请求体失败", ctx)
		return
	}
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Println("[register]json反序列化失败")
		response.FailWithDetailed("500", nil, "json反序列化失败", ctx)
		return
	}
	email := m["email"].(string)
	// 生成六位随机数
	sessionEmailCode := common.GenEmailCode(6)
	// 发送邮件
	// go common.SendEmail(email, sessionEmailCode)
	err = global.RabbitMQ.PublishQueue("email.code.wjh", email+":"+sessionEmailCode)
	if err != nil {
		log.Println("发错邮件出错了")
		return
	}
	if err != nil {
		log.Println("[register]邮箱发送验证码失败")
		response.FailWithDetailed("500", nil, "邮箱发送验证码失败", ctx)
		return
	}
	// 将数据保存到cache中
	// constant.EmailCache.Set("user_" + email, sessionEmailCode, 5 * time.Minute)
	global.RedisClient.Set(global.Context, "user_register_" + email, sessionEmailCode, 5 * time.Minute)
	response.OkWithDetailed("200", nil, "邮箱发送验证码成功", ctx)
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
	var userService service.UserService

	var tmp = struct {
		model.User
		Code string
	}{}
	// 绑定数据
	if err := ctx.Bind(&tmp); err != nil {
		log.Println("[register]注册绑定数据出错")
		response.FailWithDetailed("400", nil, "注册绑定数据出错", ctx)
		return
	}
	stu := tmp.User

	// 验证邮箱验证码是否正确
	// emailCode, found := constant.EmailCache.Get("user_" + stu.Email)
	emailCode, err := global.RedisClient.Get(global.Context, "user_register_" + stu.Email).Result()
	if err != nil || tmp.Code != emailCode {
		log.Println("[register]验证码错误")
		response.FailWithDetailed(constant.CodeError, nil, "验证码错误", ctx)
		return
	}

	// 验证字段正确性
	validate := validator.New()
	err = validate.Struct(&stu)
	if err != nil {
		log.Println("[register]字段验证错误")
		response.FailWithDetailed("400", nil, err.Error(), ctx)
		return
	}

	err, _ = userService.Register(&stu)
	// 删掉邮箱验证码
	// constant.EmailCache.Delete("stu_" + stu.Email)
	global.RedisClient.Del(global.Context, "stu_register_" + stu.Email)
	if err != nil {
		response.FailWithDetailed("400", nil, err.Error(), ctx)
		return
	}
	response.OkWithDetailed("200", nil, "注册成功", ctx)
}

func desCBCEncrypt(plainText /*明文*/, key []byte) ([]byte, error) {
	//第一步：创建des密码接口, 输入秘钥，返回接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//第二步：创建cbc分组
	// 返回一个密码分组链接模式的、底层用b解密的BlockMode接口
	// func NewCBCEncrypter(b Block, iv []byte) BlockMode
	blockSize := block.BlockSize()

	//创建一个8字节的初始化向量
	iv := bytes.Repeat([]byte("1"), blockSize)

	mode := cipher.NewCBCEncrypter(block, iv)

	//第三步：填充
	//TODO
	plainText, err = paddingNumber(plainText, blockSize)
	if err != nil {
		return nil, err
	}

	//第四步：加密
	// type BlockMode interface {
	// 	// 返回加密字节块的大小
	// 	BlockSize() int
	// 	// 加密或解密连续的数据块，src的尺寸必须是块大小的整数倍，src和dst可指向同一内存地址
	// 	CryptBlocks(dst, src []byte)
	// }

	//密文与明文共享空间，没有额外分配
	mode.CryptBlocks(plainText /*密文*/, plainText /*明文*/)

	return plainText, nil
}

//输入密文，得到明文
func desCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	//TODO
	//第一步：创建des密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//第二步：创建cbc分组
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	mode := cipher.NewCBCDecrypter(block, iv)

	//第三步：解密
	mode.CryptBlocks(encryptData /*明文*/, encryptData /*密文*/)

	//第四步: 去除填充
	//TODO
	encryptData, err = unPaddingNumber(encryptData)
	if err != nil {
		return nil, err
	}

	// return []byte("Hello world"), nil
	return encryptData, nil
}

//填充数据
func paddingNumber(src []byte, blockSize int) ([]byte, error) {

	if src == nil {
		return nil, errors.New("src长度不能小于0")
	}

	fmt.Println("调用paddingNumber")
	//1. 得到分组之后剩余的长度 5
	leftNumber := len(src) % blockSize //5

	//2. 得到需要填充的个数 8 - 5 = 3
	needNumber := blockSize - leftNumber //3

	//3. 创建一个slice，包含3个3
	b := byte(needNumber)
	newSlice := bytes.Repeat([]byte{b}, needNumber) //newSlice  ==》 []byte{3,3,3}

	fmt.Printf("newSclie : %v\n", newSlice)
	//4. 将新切片追加到src
	src = append(src, newSlice...)

	return src, nil
}

//解密后去除填充数据
func unPaddingNumber(src []byte) ([]byte, error) {
	fmt.Println("调用unPaddingNumber")
	//1. 获取最后一个字符
	lastChar := src[len(src)-1] //byte(3)

	//2. 将字符转换为数字
	num := int(lastChar) //int(3)

	//3. 截取切片(左闭右开)

	return src[:len(src)-num], nil
}

// Login 登录
func Login(ctx *gin.Context) {
	var userService service.UserService

	var loginUser model.User
	// 获取账号密码
	if err := ctx.Bind(&loginUser); err != nil {
		log.Println("[login]绑定数据出错")
		response.FailWithDetailed("400", nil, "绑定数据出错", ctx)
		return
	}
	var err error

	// todo 密码对称加密
	//key := "bosseeff"
	//plainText, err := desCBCDecrypt([]byte(loginUser.Password), []byte(key))
	//loginUser.Password=string(plainText)

	// 登录
	err, user := userService.Login(&loginUser)

	if err != nil {
		log.Println("[login]", err.Error())
		response.FailWithDetailed("400", nil, err	.Error(), ctx)
		return
	}

	// 查看用户是否是启用状态
	if user.Enable == 0 {
		log.Println("[login]该用户是禁用状态")
		response.FailWithDetailed(constant.DisableUserError, nil, "该用户是禁用状态，请联系管理员", ctx)
		return
	}

	// 登录成功，返回token
	token, err := common.ReleaseToken(*user)
	if err != nil {
		log.Println("[login]生成token出错")
		response.FailWithDetailed("500", nil, "生成token出错", ctx)
		return
	}

	// 将token存入redis，设置过期时间为6小时先
	global.RedisClient.SetEX(global.Context, "user_token_"+user.Account, token, 24 * time.Hour)

	// 返回数据
	response.OkWithDetailed("200", gin.H{"token":token, "user":dto.ToUserDtoAndRole(*user)}, "登录成功", ctx)
}

// Logout 退出登录
func Logout(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*common.Claims)
	global.RedisClient.Del(global.Context, "user_token_" + claims.Account)
	response.OkWithDetailed("200", nil, "退出登录成功", ctx)
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

func InBlackList(ctx *gin.Context){
	b,err:=model.InBlackList("201820220006")
	fmt.Println(b)
	fmt.Println(err)
}
