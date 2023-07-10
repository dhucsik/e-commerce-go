package service

import (
	"context"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/rpc"
)

type MailService struct {
	grpcClient *rpc.MailClient
}

func NewMailService(grpcClient *rpc.MailClient) *MailService {
	return &MailService{
		grpcClient: grpcClient,
	}
}

func (s *MailService) SendMail(ctx context.Context, mail *models.Mail) error {
	return s.grpcClient.SendMail(ctx, mail)
}
