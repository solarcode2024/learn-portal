package registry

import "github.com/gofiber/fiber/v2"

func Registry() (*Handler, *fiber.App, error) {
	// init infrastructure
	infrastructure := infrastructureRegistry()

	// init database
	db, err := infrastructure.Database.DatabaseConnection()
	if err != nil {
		return nil, nil, err
	}

	// init repository
	repository := repositoryRegistry(db)

	// init service
	services := serviceRegistry(repository)

	// init handler
	handler := handlerRegistry(services)

	return &handler, infrastructure.Fiber, nil
}
