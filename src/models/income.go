package models

import "time"

type Income struct {
	ID          int
	TimeOcc     string
	Description string
	Category    string
	Value       float32
	TimeAdd     time.Time
}
