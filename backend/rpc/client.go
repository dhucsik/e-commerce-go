package rpc

import (
	"context"

	"github.com/dhucsik/e-commerce-go/config"
	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MailClient struct {
	grpcClient proto.MailServiceClient
}

func NewMailClient(cfg *config.Config) (*MailClient, error) {
	conn, err := grpc.Dial("email-service:"+cfg.GRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := proto.NewMailServiceClient(conn)

	return &MailClient{
		grpcClient: client,
	}, nil
}

func (c *MailClient) SendMail(ctx context.Context, mail *models.Mail) error {
	_, err := c.grpcClient.MailSend(ctx, &proto.MailSendRequest{
		Receiver: mail.Receiver,
		Message:  mail.Message,
	})

	return err
}
