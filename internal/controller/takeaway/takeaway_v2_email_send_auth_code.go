package takeaway

import (
	"bytes"
	"context"
	"fmt"
	"gf-demo-takeaway/internal/dao"
	"gf-demo-takeaway/internal/model/do"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"gf-demo-takeaway/api/takeaway/v2"
)

// EmailData 结构体用于存储模板数据
type EmailData struct {
	VerificationCode string
}

func generateVerificationCode() string {
	code := rand.Intn(1000000) // 生成一个 6 位的随机验证码
	fmt.Println("generateVerificationCode")
	return fmt.Sprintf("%06d", code) // 格式化为 6 位数
}

func generateEmailHtml() ([]byte, string) {
	// 读取 HTML 文件
	templateContent, err := os.ReadFile("resource/template/template_authcode.html")
	if err != nil {
		log.Fatalf("Error reading HTML template file: %v", err)
	}

	// 生成验证码
	verificationCode := generateVerificationCode()

	// 解析模板
	tmpl, err := template.New("email").Parse(string(templateContent))
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// 将验证码数据传递给模板
	emailData := EmailData{VerificationCode: verificationCode}

	// 创建一个字节缓冲区来存储渲染后的 HTML 内容
	var emailHTMLContent []byte
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, emailData)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
	emailHTMLContent = buf.Bytes()

	fmt.Println("generateEmailHtml")
	return emailHTMLContent, verificationCode
}

// EmailSendAuthCode hivjkeuxaurxbbfj
func (c *ControllerV2) EmailSendAuthCode(ctx context.Context, req *v2.EmailSendAuthCodeReq) (res *v2.EmailSendAuthCodeRes, err error) {
	authcodeHtml, authcode := generateEmailHtml()

	e := &email.Email{
		To:      []string{req.Email},
		From:    "Web Exp 3 <tianxiang1014@foxmail.com>",
		Subject: "Your Authentication Code",
		HTML:    authcodeHtml,
		Headers: nil,
	}
	fmt.Println("EmailEdit")

	// 设置 SMTP 服务器配置
	smtpHost := "smtp.qq.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", "tianxiang1014@foxmail.com", "hivjkeuxaurxbbfj", smtpHost)
	fmt.Println("EmailServer")

	// 刷新验证码和邮箱的记录
	_, err = dao.Authcode.Ctx(ctx).Where(do.Authcode{Email: req.Email}).Delete()

	insertId, err := dao.Authcode.Ctx(ctx).Data(do.Authcode{
		Email:       req.Email,
		Authcode:    authcode,
		ExpiredTime: gtime.Now().Add(10 * time.Minute),
	}).InsertAndGetId()

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	res = &v2.EmailSendAuthCodeRes{
		Id: insertId,
	}

	// 发送邮件
	_ = e.Send(smtpHost+":"+smtpPort, auth)
	fmt.Println("EmailSend")

	fmt.Println("EmailGetAuthCode")
	return
}
