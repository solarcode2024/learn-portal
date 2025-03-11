package main

import (
	"learn-portal/internal/pkg/routes"
	"learn-portal/internal/registry"
)

func main() {
	handler, app, err := registry.Registry()
	if err != nil {
		panic(err)
	}

	routes.NewAPIRoutes(app, *handler)

	app.Listen(":8081")
}
