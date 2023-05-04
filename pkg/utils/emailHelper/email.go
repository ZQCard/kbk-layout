package emailHelper

import (
	"github.com/go-gomail/gomail"
)

func SendEmail(host string, port int, username, password, email, subject, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
