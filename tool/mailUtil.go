package tool

import (
	"gopkg.in/gomail.v2"
	"log"
)

func SendMail() error {
	// 1. 创建邮件消息
	msg := gomail.NewMessage()
	msg.SetHeader("From", "2368756302@qq.com")     // 发件人
	msg.SetHeader("To", "2368756302@qq.com")       // 收件人
	msg.SetHeader("Subject", "Hello from GoMail!") // 邮件主题
	msg.SetBody("text/plain", "这是一份邮件!")           // 邮件正文（纯文本）

	password, err := getEmailPassword()
	if err != nil {
		return err
	}

	// 2. 配置 SMTP 服务器
	dialer := gomail.NewDialer(
		"smtp.qq.com",       // SMTP 服务器地址（如 smtp.gmail.com）
		587,                 // SMTP 端口（Gmail 用 587）
		"2368756302@qq.com", // SMTP 用户名
		password,            // SMTP 密码（或应用专用密码）
	)

	// 3. 发送邮件
	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}

func getEmailPassword() (string, error) {
	password, err := ConfigReadEmailPassword()
	if err != nil {
		return "", err
	}
	return password, nil
}
