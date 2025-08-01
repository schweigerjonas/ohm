package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/schweigerjonas/ohm/src/db"
	"github.com/schweigerjonas/ohm/src/routes"
)

func main() {
	// Create new template engine
	engine := html.New("./src/views", ".html")

	// Pass engine to views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files (CSS, Images, ...)
	// Defines where those files can be found, e.g. for HTML files
	app.Static("/", "./src/public")

	// Initialize route handlers
	routes.Initialize(app, db)

	// Start server
	app.Listen(":3000")
}
