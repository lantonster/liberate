package service

type Service struct {
	UserService UserService
}

func NewService(
	userService UserService,
) *Service {
	return &Service{
		UserService: userService,
	}
}
