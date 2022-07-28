package models

import "time"

type Task struct {
	Created   time.Time
	Completed time.Time
	Title     string // PK
}

type Session struct {
	Started time.Time // PK
	Ended   time.Time
	Task    string // FK(Task.Title)
}

type TaskSession struct {
	Task    string    // FK(Task.Title)
	Session time.Time // FK(Session.Started)
}
