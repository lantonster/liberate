package service

import "github.com/lantonster/liberate/internal/repository"

type UserService interface {
	// Register registers a new user
	Register(username, password string) error

	// Login logs in a user
	Login(username, password string) (string, error)
}

type userService struct {
	*repository.Repo
}

func NewUserService(repo *repository.Repo) UserService {
	return &userService{Repo: repo}
}

func (s *userService) Register(username, password string) error {
	// TODO: implement
	return nil
}

func (s *userService) Login(username, password string) (string, error) {
	// TODO: implement
	return "", nil
}
