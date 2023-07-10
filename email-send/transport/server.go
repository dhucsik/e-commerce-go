package transport

import (
	"context"
	"log"
	"net"

	"github.com/dhucsik/email-send/proto"
	"google.golang.org/grpc"
)

type Server struct {
	port    string
	grpc    *grpc.Server
	handler *Handler
}

func NewServer(port string, handler *Handler) *Server {
	return &Server{
		port:    port,
		handler: handler,
	}
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	defer listener.Close()

	s.grpc = grpc.NewServer()
	proto.RegisterMailServiceServer(s.grpc, s.handler)

	go func() {
		if err := s.grpc.Serve(listener); err != nil {
			log.Fatalf("listen:%v\n", err)
		}
	}()

	<-ctx.Done()

	s.grpc.GracefulStop()

	return nil
}
