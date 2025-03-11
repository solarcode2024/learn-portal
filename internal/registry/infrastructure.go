package registry

import (
	"learn-portal/internal/pkg/database"

	"github.com/gofiber/fiber/v2"
)

type infrastructure struct {
	Database database.Database
	Fiber    *fiber.App
}

func infrastructureRegistry() infrastructure {
	app := fiber.New()

	return infrastructure{
		Database: database.NewDatabase(),
		Fiber:    app,
	}
}
