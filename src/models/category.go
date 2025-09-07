package models

import "database/sql"

type Category struct {
	ID          int
	Type        string
	Category    sql.NullString
	Subcategory string
}
