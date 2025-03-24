package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/liberate/internal/config"
	"github.com/lantonster/liberate/internal/handler"
	"github.com/lantonster/liberate/internal/router"
	"github.com/lantonster/liberate/internal/service"
)

type Server struct {
	srv *http.Server
}

func NewServer(
	cfg *config.Config,
	service *service.Service,
	handler *handler.Handler,
) *Server {
	r := gin.Default()

	router.RegisterRoutes(r, service, handler)

	return &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
			Handler: r,
		},
	}
}

func (s *Server) Start() error {
	fmt.Printf("Server started at %s\n", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
