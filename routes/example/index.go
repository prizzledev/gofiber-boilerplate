package example

import "github.com/gofiber/fiber/v2"

// SetupRoutes
// This is for the "/example" root route
func SetupRoutes(app fiber.Router) {
	// Use middleware if needed, or comment out
	// This middleware will handle everything
	app.Use(middleware)

	// define paths
	app.Get("/", indexHandler)

	// define sub routes

}

func middleware(c *fiber.Ctx) error {
	// middleware usage herr
	// e.g. c.Locals("example", "value")
	return c.Next()
}

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Example!")
}
