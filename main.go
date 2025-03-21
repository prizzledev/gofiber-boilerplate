package main

import (
	"log"

	"gofiber-boilerplate/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Creating a simple config
	// This can be modified to your specific needs
	config := fiber.Config{
		AppName:       "Routing Boilerplate by Prizzle",
		CaseSensitive: true,
	}

	// Create the fiber application using the before defined config
	app := fiber.New(config)

	// registering the default route "/" all incoming requests use
	routes.SetupRoutes(app.Group("/"))

	// start the app on port 3000
	// change this to your needs
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
