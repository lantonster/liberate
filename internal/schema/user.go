package schema

import (
	"github.com/asaskevich/govalidator"
	"github.com/lantonster/liberate/pkg/checker"
	"github.com/lantonster/liberate/pkg/errors"
	"github.com/lantonster/liberate/pkg/errors/reason"
	"github.com/lantonster/liberate/pkg/validator"
)

type SendVerificationCodeRequest struct {
	Email string `json:"email" binding:"required"`
}

type SendVerificationCodeResponse struct{}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
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

	// 检查验证码
	if len(r.Code) != 6 {
		errs = append(errs, &validator.Error{
			Field: "code",
			Error: "验证码长度不正确",
		})
		err = errors.BadRequest(reason.InvalidVerificationCode)
	}

	return errs, err
}

type RegisterResponse struct{}
