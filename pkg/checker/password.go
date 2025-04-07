package checker

import (
	"regexp"

	"github.com/lantonster/liberate/pkg/errors"
	"github.com/lantonster/liberate/pkg/errors/reason"
)

// CheckPassword 检查密码.
//
//  1. 长度在 6-20 之间
//  2. 不能有特殊字符
func CheckPassword(password string) error {
	if len(password) < 6 || len(password) > 20 {
		return errors.BadRequest(reason.PasswordLengthError).WithMsg("密码长度必须在 6-20 之间")
	}

	// 不能有特殊字符
	if _, err := regexp.MatchString(`^[a-zA-Z0-9]+$`, password); err != nil {
		return errors.BadRequest(reason.PasswordSpecialCharacterError).WithMsg("密码不能包含特殊字符").WithError(err)
	}

	return nil
}
