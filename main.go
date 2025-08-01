package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/schweigerjonas/ohm/src/routes"
)

func main() {
	// Create new template engine
	engine := html.New("./src/views", ".html")

	// Pass engine to views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files (CSS, Images, ...)
	// Defines where those files can be found, e.g. for HTML files
	app.Static("/", "./src/public")

	// Initialize route handlers
	routes.Initialize(app)

	// Start server
	app.Listen(":3000")
}
