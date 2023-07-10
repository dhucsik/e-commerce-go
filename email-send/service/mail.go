package service

import (
	"github.com/dhucsik/email-send/emailsender"
	"github.com/dhucsik/email-send/model"
)

type MailService struct {
	sender *emailsender.EmailSender
}

func NewMailService(sender *emailsender.EmailSender) *MailService {
	return &MailService{
		sender: sender,
	}
}

func (s *MailService) SendMail(mail *model.Mail) error {
	return s.sender.SendMail(mail)
}
