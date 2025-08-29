package routes

import (
	"database/sql"
	"fmt"

	"github.com/schweigerjonas/ohm/src/models"

	"github.com/gofiber/fiber/v2"
)

func getAllCategoriesHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		query := `
			SELECT * FROM categories WHERE type='expense';
		`
		var categories []models.Category

		rows, err := db.Query(query)
		if err != nil {
			return fmt.Errorf("getAllCategoriesHandler: %v", err)
		}

		defer rows.Close()

		for rows.Next() {
			var category models.Category
			if err := rows.Scan(
				&category.ID,
				&category.Type,
				&category.Category,
				&category.Subcategory,
			); err != nil {
				return fmt.Errorf("getAllCategoriesHandler: %v", err)
			}

			categories = append(categories, category)
		}

		if err := rows.Err(); err != nil {
			return fmt.Errorf("getAllCategoriesHandler: %v", err)
		}

		categoryMap := make(map[string][]string)

		for _, categoryValue := range categories {
			category := categoryValue.Category
			subcategory := categoryValue.Subcategory

			categoryMap[category] = append(categoryMap[category], subcategory)
		}

		return c.Render("select-options", fiber.Map{
			"Categories": categoryMap,
		})
	}
}
