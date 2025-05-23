// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InitializeServer() *server.Server {
	configConfig := config.LoadConfig()
	db := database.NewDB(configConfig)
	userRepo := repository.NewUserRepo(db)
	repo := repository.NewRepo(userRepo)
	userService := service.NewUserService(repo)
	serviceService := service.NewService(userService)
	userHandler := handler.NewUserHandler(serviceService)
	handlerHandler := handler.NewHandler(userHandler)
	serverServer := server.NewServer(configConfig, serviceService, handlerHandler)
	return serverServer
}

// wire.go:

var HandlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)

var ServiceSet = wire.NewSet(service.NewService, service.NewUserService)

var RepoSet = wire.NewSet(repository.NewRepo, repository.NewUserRepo)
