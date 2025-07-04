package utils

import (
	"fmt"
	"os"

	"server/internal/config"

	gomail "gopkg.in/gomail.v2"
)

func SendNotificationEmail(to string, title string, message string) error {
	plainText := message
	html := fmt.Sprintf("<p>%s</p>", message)

	return SendEmail(title, to, plainText, html)
}

func SendEmail(subject, toEmail, plainTextBody, htmlBody string) error {
	m := gomail.NewMessage()
	from := os.Getenv("USER_EMAIL")

	m.SetHeader("From", fmt.Sprintf("fitness_app <%s>", from))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", plainTextBody)
	m.AddAlternative("text/html", htmlBody)

	if err := config.MailDialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
