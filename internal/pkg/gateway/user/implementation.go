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
		slog.Error(fmt.Sprintf("gagal insert user, detail: %v\n", result.Error.Error()))
		return result.Error
	}

	return nil
}

// TODO : Implementasi repository select all
func (i *implementation) SelectAllUser() ([]domain.User, error) {
	var users []domain.User

	result := i.db.Find(&users)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("gagal select data dari db, detail: %v\n", result.Error.Error()))
		return nil, result.Error
	}

	return users, nil
}

// TODO : implementasi repository select by id
func (i *implementation) SelectOneUserByID(id int) (*domain.User, error) {
	var user domain.User

	result := i.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("gagal select data by id dari db, detail: %v\n", result.Error.Error()))
		return nil, result.Error
	}
	return &user, nil

}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &implementation{db}
}
