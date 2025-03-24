package handler

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
