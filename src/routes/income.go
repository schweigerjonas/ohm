package routes

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/schweigerjonas/ohm/src/models"

	"github.com/gofiber/fiber/v2"
)

func getAllIncomeHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := `
			SELECT * FROM income;	
		`
		var incomes []models.Income

		rows, err := db.Query(query)
		if err != nil {
			return fmt.Errorf("getAllIncomeHandler: %v", err)
		}

		defer rows.Close()

		for rows.Next() {
			var income models.Income
			if err := rows.Scan(
				&income.ID,
				&income.TimeOcc,
				&income.Description,
				&income.Category,
				&income.Value,
				&income.TimeAdd,
			); err != nil {
				return fmt.Errorf("getAllIncomeHandler: %v", err)
			}

			incomes = append(incomes, income)
		}

		if err := rows.Err(); err != nil {
			return fmt.Errorf("getAllIncomeHandler: %v", err)
		}

		return c.Render("transaction", fiber.Map{
			"Transactions": incomes,
		})
	}
}

func addIncomeHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := `
			INSERT INTO income
			(
				id, 
				time_occ, 
				description, 
				category, 
				value, 
				time_add
			) 
			VALUES (?, ?, ?, ?, ?, ?);
		`
		income := new(models.Income)

		if err := c.BodyParser(&income); err != nil {
			return fmt.Errorf("addIncomeHandler: %v", err)
		}

		result, err := db.Exec(
			query,
			nil,
			income.TimeOcc,
			income.Description,
			income.Category,
			income.Value,
			time.Now(),
		)
		if err != nil {
			return fmt.Errorf("addIncomeHandler: %v", err)
		}

		fmt.Println(result.LastInsertId())

		// Set reponse header to trigger table update
		c.Context().Response.Header.Set("HX-Trigger", "newIncome")

		return err
	}
}
