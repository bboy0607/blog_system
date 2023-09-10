package email

import (
	"crypto/tls"
	"fmt"
	"membership_system/global"

	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", to...)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}

// 發送email確認郵件
func (e *Email) SendConfirmationEmail(to []string, subject string, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", to...)

	confirmUrl := fmt.Sprintf("http://%v:%v/api/v1/users/verify_email/%v",
		global.ServerSetting.ListenAddr,
		global.ServerSetting.HttpPort,
		token,
	)
	m.SetBody("text/html", confirmUrl)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}

// 發送重置密碼郵件
func (e *Email) SendResetPasswordEmail(to []string, subject string, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", to...)

	confirmUrl := fmt.Sprintf("http://%v:%v/api/v1/users/reset_password/%v",
		global.ServerSetting.ListenAddr,
		global.ServerSetting.HttpPort,
		token,
	)
	m.SetBody("text/html", confirmUrl)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
