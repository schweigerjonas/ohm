package routes

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/schweigerjonas/ohm/src/models"

	"github.com/gofiber/fiber/v2"
)

func getAllExpensesHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := `
			SELECT * FROM expenses;	
		`
		var expenses []models.Expense

		rows, err := db.Query(query)
		if err != nil {
			return fmt.Errorf("getAllExpensesHandler: %v", err)
		}

		defer rows.Close()

		for rows.Next() {
			var expense models.Expense
			if err := rows.Scan(&expense.ID, &expense.TimeOcc, &expense.Description, &expense.Category, &expense.Value, &expense.TimeAdd); err != nil {
				return fmt.Errorf("getAllExpensesHandler: %v", err)
			}

			expenses = append(expenses, expense)
		}

		if err := rows.Err(); err != nil {
			return fmt.Errorf("getAllExpensesHandler: %v", err)
		}

		return c.Render("expense", fiber.Map{
			"Expenses": expenses,
		})
	}
}

func addExpenseHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := `
			INSERT INTO expenses 
			(
				id, 
				time_occ, 
				description, 
				category, 
				value, 
				time_add
			) 
			VALUES (?, ?, ?, ?, ?, ?);"
		`
		expense := new(models.Expense)

		if err := c.BodyParser(&expense); err != nil {
			return fmt.Errorf("addExpenseHandler: %v", err)
		}

		result, err := db.Exec(
			query,
			nil,
			expense.TimeOcc,
			expense.Description,
			expense.Category,
			expense.Value,
			time.Now(),
		)
		if err != nil {
			return fmt.Errorf("addExpenseHandler: %v", err)
		}

		fmt.Print(result.LastInsertId())

		return err
	}
}
