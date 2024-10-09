package mailer

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendContactEmail(name, email, comment string) error {
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_PASSWORD")

	auth := smtp.PlainAuth("", senderEmail, password, smtpHost)

	toCompany := "company@example.com"
	subjectCompany := "New contact from " + name
	bodyCompany := fmt.Sprintf("Name: %s\nEmail: %s\nComment: %s", name, email, comment)
	messageCompany := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", senderEmail, toCompany, subjectCompany, bodyCompany)

	subjectUser := "Thank you for contacting us"
	bodyUser := fmt.Sprintf("Hi %s,\n\nThank you for your message! We will get back to you soon.\n\nComment: %s", name, comment)
	messageUser := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", senderEmail, email, subjectUser, bodyUser)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{toCompany}, []byte(messageCompany))
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{email}, []byte(messageUser))
	if err != nil {
		return err
	}

	return nil
}
