package schema

import (
	"github.com/asaskevich/govalidator"
	"github.com/lantonster/liberate/pkg/checker"
	"github.com/lantonster/liberate/pkg/errors"
	"github.com/lantonster/liberate/pkg/errors/reason"
	"github.com/lantonster/liberate/pkg/validator"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *RegisterRequest) Check() (errs validator.ValidationErrors, err error) {
	// 检查密码
	if e := checker.CheckPassword(r.Password); e != nil {
		errs = append(errs, &validator.Error{
			Field: "password",
			Error: e.Error(),
		})
		err = errors.BadRequest(reason.PasswordLengthError)
	}

	// 检查邮箱
	if !govalidator.IsEmail(r.Email) {
		errs = append(errs, &validator.Error{
			Field: "email",
			Error: "邮箱格式不正确",
		})
		err = errors.BadRequest(reason.EmailInvalid)
	}

	return errs, err
}

type RegisterResponse struct{}
