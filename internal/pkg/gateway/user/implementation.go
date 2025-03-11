package user

import (
	"fmt"
	"learn-portal/internal/domain"
	"log/slog"

	"gorm.io/gorm"
)

// apapun yang repository butuhkan
type implementation struct {
	db *gorm.DB
}

func (i *implementation) InsertUser(data *domain.User) error {
	// insert into users (name, phone) VALUES ($1, $2)
	result := i.db.Create(data)
	if result.Error != nil {
		slog.Error(fmt.Sprintln("gagal insert user, detail: %v", result.Error.Error()))
		return result.Error
	}

	return nil
}

// TODO : Implementasi repository select all
func (i *implementation) SelectAllUser() ([]domain.User, error) {
	panic("unimplemented")
}

// TODO : implementasi repository select by id
func (i *implementation) SelectOneUserByID() (*domain.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &implementation{db}
}
