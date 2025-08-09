package models

import "time"

type Expense struct {
	// ID          int
	TimeOcc     time.Time
	Description string
	Category    string
	Value       float32
	// TimeAdd     time.Time
}
