package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dhucsik/e-commerce-go/config"
	"github.com/dhucsik/e-commerce-go/transport/http/handler"
	middleware2 "github.com/dhucsik/e-commerce-go/transport/http/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Auth interface {
	SignInMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

type Server struct {
	cfg     *config.Config
	handler *handler.Manager
	App     *echo.Echo
	auth    Auth
}

func NewServer(cfg *config.Config, handler *handler.Manager) *Server {
	auth := middleware2.NewJWTAuth(cfg)
	return &Server{
		cfg:     cfg,
		handler: handler,
		auth:    auth,
	}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()

	go func() {
		if err := s.App.Start(fmt.Sprintf(":%s", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%v\n", err)
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%v", err)
	}

	return nil

}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	return e
}
