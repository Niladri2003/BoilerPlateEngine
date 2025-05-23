// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/niladri2003/BoilerPlateEngine/db"
	engine2 "github.com/niladri2003/BoilerPlateEngine/engine"
	"github.com/niladri2003/BoilerPlateEngine/routes"
	"log"
)

func main() {
	// Initialize the boilerplate engine
	engine, err := engine2.NewBoilerplateEngine()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Fiber app
	app := fiber.New()
	app.Use(logger.New())
	//
	if err := db.ConnectToMongo(); err != nil {
		log.Fatal(err)
	}

	// Register routes
	routes.RegisterRoutes(app, engine)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
