package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App, db *sql.DB) {
	// home page handlers
	app.Get("/", homeHandler(db))

	// expense handlers
	app.Get("/api/expense", getAllExpensesHandler(db))
	app.Post("/api/expense", addExpenseHandler(db))

	// income handlers

	// category handlers
	app.Get("api/category/expense", getAllExpenseCategoriesHandler(db))
}
