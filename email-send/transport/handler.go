package transport

import (
	"context"

	"github.com/dhucsik/email-send/model"
	"github.com/dhucsik/email-send/proto"
	"github.com/dhucsik/email-send/service"
)

type Handler struct {
	proto.UnimplementedMailServiceServer
	srv *service.Manager
}

func NewHandler(srv *service.Manager) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) MailSend(ctx context.Context, req *proto.MailSendRequest) (*proto.Empty, error) {
	mail := &model.Mail{
		Receiver: req.Receiver,
		Message:  req.Message,
	}

	if err := h.srv.Mail.SendMail(mail); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
