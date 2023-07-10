package service

import (
	"errors"

	"github.com/dhucsik/email-send/emailsender"
	"github.com/dhucsik/email-send/model"
)

type IMailService interface {
	SendMail(mail *model.Mail) error
}

type Manager struct {
	Mail IMailService
}

func NewManager(sender *emailsender.EmailSender) (*Manager, error) {
	if sender == nil {
		return nil, errors.New("no email sender provided")
	}

	mailSrv := NewMailService(sender)

	return &Manager{
		Mail: mailSrv,
	}, nil
}
