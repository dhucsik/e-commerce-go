package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/dhucsik/e-commerce-go/config"
	"github.com/dhucsik/e-commerce-go/rpc"
	"github.com/dhucsik/e-commerce-go/service"
	"github.com/dhucsik/e-commerce-go/storage"
	"github.com/dhucsik/e-commerce-go/transport/http"
	"github.com/dhucsik/e-commerce-go/transport/http/handler"
)

func main() {
	log.Fatalln(run())
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		return err
	}

	stg, err := storage.New(ctx, conf)
	if err != nil {
		return err
	}

	grpc, err := rpc.NewMailClient(conf)
	if err != nil {
		return err
	}

	svc, err := service.NewManager(stg, grpc)
	if err != nil {
		return err
	}

	h := handler.NewManager(svc)
	HTTPServer := http.NewServer(conf, h)

	return HTTPServer.StartHTTPServer(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
