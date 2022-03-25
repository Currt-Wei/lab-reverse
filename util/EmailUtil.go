package util
//
import (
	"fmt"
	"github.com/beego/beego/v2/adapter/utils"
	"lab-reverse/constant"
	"math/rand"
	"strings"
	"time"
)

func GenEmailCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}

func SendEmail(to_email, msg string) error {
	emailSetting := constant.EmailSetting

	emailConfig := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%s}`,
		emailSetting.Username,
		emailSetting.Password,
		emailSetting.Host,
		emailSetting.Port)
	emailConn := utils.NewEMail(emailConfig)  // beego下的
	emailConn.From = strings.TrimSpace(emailSetting.Username)
	emailConn.To = []string{strings.TrimSpace(to_email)}
	emailConn.Subject = "智能系统未来创新实验室"
	//注意这里我们发送给用户的是激活请求地址
	emailConn.Text = msg
	err := emailConn.Send()
	return err
}