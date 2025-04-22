package tool

import (
	"gopkg.in/gomail.v2"
	"log"
)

var from, to, password string

func MailInit() error {
	fromEmailUser, err := ConfigReadFromEmailUser()
	if err != nil {
		return err
	}
	from = fromEmailUser
	toEmailUser, err := ConfigReadToEmailUser()
	if err != nil {
		return err
	}
	to = toEmailUser
	pw, err := ConfigReadEmailPassword()
	if err != nil {
		return err
	}
	password = pw
	return nil
}
func SendMail(text string) error {
	err := MailInit()
	if err != nil {
		return err
	}
	// 1. 创建邮件消息
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)      // 发件人
	msg.SetHeader("To", to)          // 收件人
	msg.SetHeader("Subject", "资产警告") // 邮件主题
	msg.SetBody("text/plain", text)  // 邮件正文（纯文本）

	// 2. 配置 SMTP 服务器
	dialer := gomail.NewDialer(
		"smtp.qq.com", // SMTP 服务器地址（如 smtp.gmail.com）
		587,           // SMTP 端口（Gmail 用 587）
		from,          // SMTP 用户名
		password,      // SMTP 密码（或应用专用密码）
	)

	// 3. 发送邮件
	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
