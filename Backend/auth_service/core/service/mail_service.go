package service

import (
	"github.com/wneessen/go-mail"
)

type MailService struct {
	senderMail     string
	senderPassword string
}

func NewMailService(senderMail string, senderPassword string) *MailService {
	return &MailService{senderMail: senderMail, senderPassword: senderPassword}
}

func (s *MailService) sendEmail(recipient, subject, body string) error {
	message := mail.NewMsg()

	err := message.From(s.senderMail)
	if err != nil {
		return err
	}

	err = message.To(recipient)
	if err != nil {
		return err
	}

	message.Subject(subject)

	message.SetBodyString(mail.TypeTextPlain, body)
	if err != nil {
		return err
	}

	// Create a new SMTP dialer
	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithPort(587),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(s.senderMail), mail.WithPassword(s.senderPassword))
	if err != nil {
		return err
	}

	if err := client.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
