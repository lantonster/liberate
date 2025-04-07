package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lantonster/liberate/pkg/errors"
	"github.com/lantonster/liberate/pkg/errors/reason"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const passwordCost = 16

type User struct {
	Id        int64          `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Email    string `json:"email"`    // 邮箱
	Uuid     string `json:"uuid"`     // 用户唯一标识
}

// NewUser 根据邮箱和密码生成用户对象
func NewUser(email, password string) (*User, error) {
	// 生成用户唯一标识
	uuid := uuid.New().String()

	// 生成用户名
	username := uuid[:8]

	// 生成密码
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return nil, errors.InternalServer(reason.GeneratePasswordHashFailed).WithMsg("生成密码哈希失败").WithError(err)
	}

	return &User{
		Username: username,
		Email:    email,
		Password: string(passwordHash),
		Uuid:     uuid,
	}, nil
}
