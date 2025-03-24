//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lantonster/liberate/internal/config"
	"github.com/lantonster/liberate/internal/server"
)

func InitializeServer() *server.Server {
	wire.Build(
		config.LoadConfig,
		server.NewServer,
	)
	return &server.Server{}
}
