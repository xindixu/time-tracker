package db

import (
	"github.com/boltdb/bolt"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
)

func InitDB() error {

	db, err := bolt.Open("todo-time-tracker.db", 0600, nil)
	if err != nil {
		return err
	}

	err = taskDB.Setup(db)
	return err
}
