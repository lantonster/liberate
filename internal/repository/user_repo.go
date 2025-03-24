package repository

import (
	"context"

	"github.com/lantonster/liberate/internal/model"
	"gorm.io/gorm"
)

// UserRepo 定义用户数据访问接口
type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
}

// userRepo 实现 UserRepo 接口
type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	// TODO: 实现创建逻辑
	return nil
}
