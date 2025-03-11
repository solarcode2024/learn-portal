package registry

import (
	"learn-portal/internal/domain"
	"learn-portal/internal/service/user"
)

type services struct {
	UserService domain.UserService
}

func serviceRegistry(repository repositories) services {
	return services{
		UserService: user.NewUserService(repository.UserRepository),
	}
}
