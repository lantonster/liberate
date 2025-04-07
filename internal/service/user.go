package service

import (
	"context"

	"github.com/lantonster/liberate/internal/model"
	"github.com/lantonster/liberate/internal/repository"
	"github.com/lantonster/liberate/pkg/log"
)

type UserService interface {
	// CheckEmailExists 检查邮箱是否存在
	CheckEmailExists(c context.Context, email string) (bool, error)

	// Register 注册用户
	Register(c context.Context, email, password string) error

	// Login 登录用户
	Login(username, password string) (string, error)
}

type userService struct {
	*repository.Repo
}

func NewUserService(repo *repository.Repo) UserService {
	return &userService{Repo: repo}
}

func (s *userService) CheckEmailExists(c context.Context, email string) (bool, error) {
	user, err := s.UserRepo.GetByEmail(c, email)
	if err != nil {
		return false, err
	}
	return user.Email == email, nil
}

func (s *userService) Register(c context.Context, email, password string) error {
	// 生成用户
	user, err := model.NewUser(email, password)
	if err != nil {
		log.WithContext(c).Errorf("failed to generate user: %v", err)
		return err
	}

	// 创建用户
	if err := s.UserRepo.Create(c, user); err != nil {
		log.WithContext(c).Errorf("failed to create user: %v", err)
		return err
	}
	return nil
}

func (s *userService) Login(username, password string) (string, error) {
	// TODO: implement
	return "", nil
}
