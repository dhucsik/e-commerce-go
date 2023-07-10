package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/dhucsik/email-send/config"
	"github.com/dhucsik/email-send/emailsender"
	"github.com/dhucsik/email-send/service"
	"github.com/dhucsik/email-send/transport"
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
	log.Println(conf)
	sender := emailsender.New(ctx, conf)
	log.Println("sender")
	svc, err := service.NewManager(sender)
	if err != nil {
		return err
	}
	log.Println("svc")
	h := transport.NewHandler(svc)
	log.Println("handler")
	server := transport.NewServer(conf.ListenerPort, h)
	log.Println("server")
	return server.Start(ctx)
}

func gracefullyShutdown(cancel context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)

	go func() {
		log.Print(<-osC)
		cancel()
	}()
}
