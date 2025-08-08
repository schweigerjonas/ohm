package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App, db *sql.DB) {
	app.Get("/", homeHandler(db))
	app.Get("/api/expense", expenseHandler(db))
}
