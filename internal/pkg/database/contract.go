package database

import "gorm.io/gorm"

type Database interface {
	DatabaseConnection() (*gorm.DB, error)
}
