package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/liberate/internal/service"
)

type UserHandler struct {
	*service.Service
}

func NewUserHandler(service *service.Service) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	// 实现注册账号的逻辑
}
