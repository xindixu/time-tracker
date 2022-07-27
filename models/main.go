package models

import "time"

type Task struct {
	ID        int
	Created   time.Time
	Completed time.Time
	Title     string
}
