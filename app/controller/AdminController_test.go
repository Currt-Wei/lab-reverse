package controller
//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	utils "github.com/Valiben/gin_unit_test"
//	"testing"
//)
//
//func init(){
//	router :=gin.Default()
//	router.GET("/findAllUser", FindAllUser)
//
//}
//
//type Response struct {
//	Errno  string `json:"errno"`
//	Errmsg string `json:"errmsg"`
//}
//
//func TestFindAllUser(t *testing.T) {
//	resp:=Response{}
//	err := utils.TestHandlerUnMarshalResp("GET", "/login", "json", user, &resp)
//	if err != nil {
//		t.Errorf("TestLoginHandler: %v\n", err)
//		return
//	}
//	// 得到返回数据结构体， 至此，完美完成一次post请求测试，
//	// 如果需要benchmark 输出性能报告也是可以的
//	fmt.Println("result:", resp)
//}
