package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lab-reverse/app/middleware"
	"lab-reverse/app/model"
	"lab-reverse/app/service"
	"lab-reverse/constant"
	"net/http"
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

// Login 登陆
func Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.LoginFail,
			"msg":    "登录失败",
			"data":   err.Error(),
		})
		return
	}

	//TODO 查找数据库
	user, err := service.FindUserByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.LoginFail,
			"msg":    err.Error(),
			"data":   "登录失败",
		})
		return
	}

	// 密码错误
	if u.Password != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.LoginFail,
			"msg":    "登录失败",
			"data":   err.Error(),
		})
		return
	}

	token := generateToken(c, *user)
	var role string
	// 用户角色
	//if len(user.Role) == 0 {
	//	role = "user"
	//} else {
	//	role = user.Role[0].RoleName
	//}
	c.JSON(http.StatusOK, gin.H{
		"status": constant.LoginSuccess,
		"msg":    "登陆成功",
		"data": gin.H{
			"token": token,
			"role":  role,
		},
	})

	return
}

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

