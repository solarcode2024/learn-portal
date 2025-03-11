package registry

import (
	"learn-portal/internal/domain"
	"learn-portal/internal/pkg/gateway/user"

	"gorm.io/gorm"
)

type repositories struct {
	UserRepository domain.UserRepository
}

func repositoryRegistry(db *gorm.DB) repositories {
	return repositories{
		UserRepository: user.NewUserRepository(db),
	}
}
