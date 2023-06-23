package app

import (
	"github.com/gofiber/fiber"
	"go-practice/backendApi/routes"
)
func AppInit() {
	app := fiber.New()

	// Routes
	routes.HelloRoute(app)

	// start server
	app.Listen(3000)

}

