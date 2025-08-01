package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App /*, db *sql.DB*/) {
	app.Get("/", homeHandler( /*db*/ ))
}
