package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lantonster/liberate/pkg/errors"
	"github.com/lantonster/liberate/pkg/errors/reason"
	"github.com/lantonster/liberate/pkg/log"
	"github.com/lantonster/liberate/pkg/validator"
)

type Handler struct {
	UserHandler *UserHandler
}

func NewHandler(
	userHandler *UserHandler,
) *Handler {
	return &Handler{
		UserHandler: userHandler,
	}
}

// BindAndValidate binds and validates the request body
func BindAndValidate(c *gin.Context, req any) (validator.ValidationErrors, error) {
	if err := c.ShouldBind(req); err != nil {
		err = errors.BadRequest(reason.RequestFormatError).WithMsg(err.Error()).WithError(err)
		return nil, err
	}

	fields, err := validator.Check(c, req)
	if err != nil {
		log.WithContext(c).With("url", c.Request.URL.Path).Errorf("failed to validate request: %v", err)
		return fields, err
	}

	return fields, nil
}
