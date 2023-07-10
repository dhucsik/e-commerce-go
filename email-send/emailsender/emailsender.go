package emailsender

import (
	"context"
	"net/smtp"

	"github.com/dhucsik/email-send/config"
	"github.com/dhucsik/email-send/model"
)

type EmailSender struct {
	host   string
	port   string
	sender string
	auth   smtp.Auth
}

func New(ctx context.Context, conf *config.Config) *EmailSender {
	auth := smtp.PlainAuth("", conf.SenderEmail, conf.Password, conf.SmtpHost)
	return &EmailSender{
		host:   conf.SmtpHost,
		port:   conf.SmtpPort,
		sender: conf.SenderEmail,
		auth:   auth,
	}
}

func (es *EmailSender) SendMail(mail *model.Mail) error {
	err := smtp.SendMail(es.host+":"+es.port, es.auth, es.sender, []string{mail.Receiver}, []byte(mail.Message))
	if err != nil {
		return err
	}

	return nil
}
