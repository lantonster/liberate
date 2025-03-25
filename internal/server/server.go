package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/liberate/docs/api"
	"github.com/lantonster/liberate/internal/config"
	"github.com/lantonster/liberate/internal/handler"
	"github.com/lantonster/liberate/internal/router"
	"github.com/lantonster/liberate/internal/service"
	"github.com/lantonster/liberate/pkg/color"
	"github.com/lantonster/liberate/pkg/log"
)

type Server struct {
	srv *http.Server
}

func NewServer(
	conf *config.Config,
	service *service.Service,
	handler *handler.Handler,
) *Server {
	r := gin.Default()

	router.RegisterRoutes(r, service, handler)

	api.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", conf.Server.Port)
	info := color.Green.Sprintf("API docs: http://localhost:%d/swagger/index.html", conf.Server.Port)
	defer log.WithContext(context.Background()).Infof(info)

	return &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", conf.Server.Port),
			Handler: r,
		},
	}
}

func (s *Server) Start() error {
	log.WithContext(context.Background()).Info(color.Green.Sprintf("Server started at %s", s.srv.Addr))
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
