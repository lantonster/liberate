package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lantonster/liberate/internal/config"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg *config.Config) *Server {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

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
