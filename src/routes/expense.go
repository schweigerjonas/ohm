package routes

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/schweigerjonas/ohm/src/models"

	"github.com/gofiber/fiber/v2"
)

func expenseHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := "INSERT INTO expenses (time_occ, description, category, value, time_add) VALUES (?, ?, ?, ?, ?);"

		expense := new(models.Expense)

		if err := c.BodyParser(&expense); err != nil {
			return fmt.Errorf("BodyParser in addExpenseHandler: %v", err)
		}

		result, err := db.Exec(query, expense.TimeOcc, expense.Description, expense.Category, expense.Value, time.Now())
		if err != nil {
			return fmt.Errorf("addExpenseHandler: %v", err)
		}

		fmt.Print(result.LastInsertId())

		return c.Render("index", fiber.Map{
			"Title":   "Ohm",
			"Message": "Welcome to Budget App after Insert.",
		})
	}
}
