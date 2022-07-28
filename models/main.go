package models

import "time"

type Task struct {
	Created   time.Time
	Completed time.Time
	Title     string // PK
}
