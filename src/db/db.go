package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "database/budget.db"
const createIncome string = `
	CREATE TABLE IF NOT EXISTS income (
		id INTEGER NOT NULL PRIMARY KEY,
		time_occ DATETIME NOT NULL,
		description TEXT,
		category TEXT,
		value REAL,
		time_add DATETIME NOT NULL
	);`
const createExpense string = `
	CREATE TABLE IF NOT EXISTS expenses (	
		id INTEGER NOT NULL PRIMARY KEY,
		time_occ DATETIME NOT NULL,
		description TEXT,
		category TEXT,
		value REAL,
		time_add DATETIME NOT NULL
	);`

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(createIncome)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(createExpense)
	if err != nil {
		return nil, err
	}

	return db, nil
}
