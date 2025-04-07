package repository

import (
	"context"

	"github.com/lantonster/liberate/internal/model"
	"github.com/lantonster/liberate/pkg/orm"
	"gorm.io/gorm"
)

// UserRepo 定义用户数据访问接口
type UserRepo interface {
	// Create 创建用户
	Create(ctx context.Context, user *model.User) error

	// GetByEmail 根据邮箱查询用户
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

// userRepo 实现 UserRepo 接口
type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (user model.User, err error) {
	err = r.db.WithContext(ctx).Where(orm.Q.User.Email.Eq(email)).Find(&user).Error
	return
}
