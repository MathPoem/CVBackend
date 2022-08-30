package mymailer

import (
	"net/smtp"
)

type mailer struct {
	from string
	pwd  string
	srv  string
	port string
}

func NewMailer(pwd string) *mailer {
	return &mailer{
		from: "ecclesiastes.solomon@yandex.ru",
		pwd:  pwd,
		srv:  "smtp.yandex.ru",
		port: ":587",
	}
}

type IMailer interface {
	SendMail(email string, key string) error
}

type mailerService struct {
	Mail IMailer
}

func NewMailerService(pwd string) *mailerService {
	return &mailerService{
		Mail: NewMailer(pwd),
	}
}

func (m *mailer) SendMail(email string, key string) error {
	// Set up authentication information.

	smtpServer := m.srv
	auth := smtp.PlainAuth(
		"",
		m.from,
		m.pwd,
		smtpServer,
	)

	body := "key ----> " + key

	body += "\r\nэто ключ для авторизации на сервисе campusview.info"

	err := smtp.SendMail(
		smtpServer+m.port,
		auth,
		m.from,
		[]string{email},
		[]byte(body),
	)
	if err != nil {
		return err
	}
	return nil
}
