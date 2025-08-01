package routes

import (
	"github.com/gofiber/fiber/v2"
)

func homeHandler( /*db *sql.DB*/ ) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":   "Ohm",
			"Message": "Welcome to Budget App!",
		})
	}
}
