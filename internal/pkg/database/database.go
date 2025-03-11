package database

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	PORT     = "8080"
	DB_NAME  = "portal"
	USER     = "user"
	PASSWORD = "password"
)

type implementation struct{}

func (i *implementation) DatabaseConnection() (*gorm.DB, error) {
	slog.Info("CONNECTING TO DATABASE INSTANCE!")

	dsn := fmt.Sprintf(
		"user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
		USER,
		PASSWORD,
		DB_NAME,
		HOST,
		PORT,
	)
	driver := postgres.Open(dsn)
	return gorm.Open(driver, &gorm.Config{})
}

func NewDatabase() Database {
	return &implementation{}
}
