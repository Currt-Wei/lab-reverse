package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lab-reverse/app/common"
	"lab-reverse/constant"
	"lab-reverse/global"
	"net/http"
	"strings"
	"time"
)

// 定义一个jwt对象
type JWT struct {
	// 声明签名信息
	SigningKey []byte
}

// 初始化jwt对象
func NewJWT() *JWT {
	return &JWT{
		[]byte("实验室预约"),
	}
}

// 自定义有效载荷(这里采用自定义的Name和Email作为有效载荷的一部分)
type CustomClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// StandardClaims结构体实现了Claims接口(Valid()函数)
	jwt.StandardClaims
}

// 调用jwt-go库生成token
//指定编码的算法为 jwt.SigningMethodHS256
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#Token
	//返回一个token的结构体指针
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// token解码
func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#ParseWithClaims
	//输入用户自定义的Claims结构体对象,token,以及自定义函数来解析token字符串为jwt的Token结构体指针
	//Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	//func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		// https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		//jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token不可用")
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token过期")
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("无效的token")
			} else {
				return nil, fmt.Errorf("token不可用")
			}
		}
	}

	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token无效")
}

// 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		// Authorization:Bearer xxxxx
		tokenString := ctx.GetHeader("Authorization")

		// 验证token的格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"data": nil,
				"msg": "token格式出错",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)

		// 该用户已经被禁用
		if claims.Enable == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": constant.DisableUserError,
				"data": nil,
				"msg": "该用户是禁用状态，请联系管理员",
			})
			ctx.Abort()
			return
		}

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"data": nil,
				"msg": "token出错",
			})
			ctx.Abort()
			return
		}

		// 判断该用户的token是否一致
		if tokenString != global.RedisClient.Get(global.Context, "user_token_"+claims.Account).Val() {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"data": nil,
				"msg": "token出错",
			})
			ctx.Abort()
			return
		}

		// 给该用户续期
		global.RedisClient.SetEX(global.Context, "user_token_"+claims.Account, tokenString, 6 * time.Hour)

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
