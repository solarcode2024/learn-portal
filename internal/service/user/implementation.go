package user

import (
	"fmt"
	"learn-portal/internal/domain"
	"log/slog"
)

// apapun yang dibutuhkan oleh NewUserService()
type implementation struct {
	repository domain.UserRepository
}

func (i *implementation) GetAllUser() ([]domain.User, error) {
	users, err := i.repository.SelectAllUser()
	if err != nil {
		slog.Error(fmt.Sprintf("gagal mengambil list users, detail %v", err.Error()))
		return nil, err
	}

	return users, nil
}

// TODO : PR implementasi oleh afif
func (i *implementation) GetUserProfile(id int) (*domain.User, error) {
	user, err := i.repository.SelectOneUserByID(id)
	if err != nil {
		slog.Error(fmt.Sprintf("gagal mengambil data user, detail %v", err.Error()))
		return nil, err
	}
	return user, nil
}

func (i *implementation) Register(data *domain.User) error {
	// permisalan: validasi

	if err := i.repository.InsertUser(data); err != nil {
		slog.Error(fmt.Sprintf("gagal menambahkan user, detail %v", err.Error()))
		return err
	}

	// permisalan: generate otp

	// permisalan: kirim email

	return nil
}

func NewUserService(repository domain.UserRepository) domain.UserService {
	return &implementation{repository}
}
