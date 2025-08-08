package routes

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func expenseHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := "INSERT INTO expenses (time_occ, description, category, value, time_add) VALUES (?, ?, ?, ?, ?);"

		rows, err := db.Exec(query, "test", "test", "test", 10.0, "test")
		if err != nil {
			return err
		}
		fmt.Print(rows)

		return c.Render("index", fiber.Map{
			"Title":   "Ohm",
			"Message": "Welcome to Budget App after Insert.",
		})
	}
}
