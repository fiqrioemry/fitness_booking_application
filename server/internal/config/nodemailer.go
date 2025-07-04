package config

import (
	"log"
	"os"

	gomail "gopkg.in/gomail.v2"
)

var MailDialer *gomail.Dialer

func InitMailer() {
	email := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASSWORD")

	port := 587
	MailDialer = gomail.NewDialer("smtp.gmail.com", port, email, password)
	MailDialer.TLSConfig = nil
	log.Println("Mailer configured")
}
