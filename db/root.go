package db

import (
	"github.com/boltdb/bolt"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	models "github.com/xindixu/todo-time-tracker/models"
)

func InitDB() error {
	var err error
	models.TTTDB, err = bolt.Open("todo-time-tracker.db", 0600, nil)
	if err != nil {
		return err
	}

	err = taskDB.Setup()
	return err
}
