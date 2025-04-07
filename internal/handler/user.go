package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/liberate/internal/schema"
	"github.com/lantonster/liberate/internal/service"
	"github.com/lantonster/liberate/pkg/errors"
	"github.com/lantonster/liberate/pkg/errors/reason"
	"github.com/lantonster/liberate/pkg/resp"
)

type UserHandler struct {
	*service.Service
}

func NewUserHandler(service *service.Service) *UserHandler {
	return &UserHandler{Service: service}
}

// Register godoc
//
//	@Summary		Register
//	@Description	Register a new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			req	body		schema.RegisterRequest	true	"Register request"
//	@Success		200	{object}	schema.RegisterResponse
//	@Router			/users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var (
		req = &schema.RegisterRequest{}
		res = &schema.RegisterResponse{}
		err error
	)

	// 绑定请求参数并验证
	if errs, err := BindAndValidate(c, req); err != nil {
		resp.Response(c, err, errs)
		return
	}

	// 检查邮箱是否存在
	exists, err := h.UserService.CheckEmailExists(c, req.Email)
	if err != nil {
		resp.Response(c, err, nil)
		return
	} else if exists {
		resp.Response(c, errors.BadRequest(reason.EmailExists), nil)
		return
	}

	// 注册用户
	err = h.UserService.Register(c, req.Email, req.Password)
	resp.Response(c, err, res)
}
