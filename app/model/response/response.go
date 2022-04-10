package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应体
type Response struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code string, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func OkWithDetailed(code string, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

func FailWithDetailed(code string, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

