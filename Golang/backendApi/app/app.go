package app

import (
	"github.com/gofiber/fiber"
	"go-practice/backendApi/routes"
	"go-practice/backendApi/databases"
	"log"
)
func AppInit() {
	app := fiber.New()

	DbInit()

	// Routes
	routes.HelloRoute(app)

	// start server
	app.Listen(3000)

}

func DbInit() {
	_, dbError := databases.CockroachInit()
	if dbError != nil {
		log.Fatal(dbError)
	}
	log.Println("Database Initialized")
}


func RoutesInit(){
	
}