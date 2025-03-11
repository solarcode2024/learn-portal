package registry

import "learn-portal/internal/handler/user"

type Handler struct {
	User user.Handler
}

func handlerRegistry(service services) Handler {
	return Handler{
		User: user.NewUserHandler(service.UserService),
	}
}
