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
//	@Router			/user/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var (
		req = &schema.RegisterRequest{}
		res = &schema.RegisterResponse{}
		err error
	)

	if err = c.ShouldBind(req); err != nil {
		err = errors.BadRequest(reason.RequestFormatError).WithMsg(err.Error()).WithError(err)
		resp.Response(c, err, nil)
		return
	}

	err = h.UserService.Register(c, req.Email, req.Password)
	resp.Response(c, err, res)
}
