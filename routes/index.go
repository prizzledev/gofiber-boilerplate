package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiber-boilerplate/routes/example"
)

// SetupRoutes
// This is for the "/" root route
func SetupRoutes(app fiber.Router) {
	// Use middleware if needed, or comment out
	// This middleware will handle everything
	app.Use(middleware)

	// define paths
	app.Get("/", indexHandler)

	// define sub routes
	// even thought the directory path is /example, we can name it different, but
	// it's not recommended
	example.SetupRoutes(app.Group("/example"))
}

func middleware(c *fiber.Ctx) error {
	// middleware usage herr
	// e.g. c.Locals("example", "value")
	return c.Next()
}

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome!")
}
