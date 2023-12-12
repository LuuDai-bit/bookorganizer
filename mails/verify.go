package mails

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

type VerifyMail struct{}

func (s *VerifyMail) SendVerifyCode(destinationEmail string, verifyCode string) {
	to := []string{
		destinationEmail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	template, err := template.ParseFiles("templates/verify_code.html")
	if err != nil {
		panic(err)
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Verify code for Book Organizer \n%s\n\n", mimeHeaders)))

	template.Execute(&body, struct {
		VerifyCode string
	}{
		VerifyCode: verifyCode,
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		panic(err)
	}
}
