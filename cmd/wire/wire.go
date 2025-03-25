//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/lantonster/liberate/internal/config"
	"github.com/lantonster/liberate/internal/database"
	"github.com/lantonster/liberate/internal/handler"
	"github.com/lantonster/liberate/internal/repository"
	"github.com/lantonster/liberate/internal/server"
	"github.com/lantonster/liberate/internal/service"
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var RepoSet = wire.NewSet(
	repository.NewRepo,
	repository.NewUserRepo,
)

func InitializeServer() *server.Server {
	wire.Build(
		config.LoadConfig,
		server.NewServer,
		database.NewDB,
		HandlerSet,
		ServiceSet,
		RepoSet,
	)
	return &server.Server{}
}
