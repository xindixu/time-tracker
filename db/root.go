package db

import (
	"github.com/boltdb/bolt"
	sessionDB "github.com/xindixu/todo-time-tracker/db/sessions"
	taskSessionDB "github.com/xindixu/todo-time-tracker/db/task-sessions"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	m "github.com/xindixu/todo-time-tracker/models"
)

func InitDB() error {
	var err error
	m.TTTDB, err = bolt.Open("todo-time-tracker.db", 0600, nil)
	if err != nil {
		return err
	}

	err = taskDB.Setup()
	if err != nil {
		return err
	}

	err = sessionDB.Setup()
	if err != nil {
		return err
	}

	err = taskSessionDB.Setup()
	if err != nil {
		return err
	}
	return err
}
