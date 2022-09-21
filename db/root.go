package db

import (
	"github.com/boltdb/bolt"
	sessionDB "github.com/xindixu/todo-time-tracker/db/sessions"
	taskSessionDB "github.com/xindixu/todo-time-tracker/db/task-sessions"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	m "github.com/xindixu/todo-time-tracker/models"
	"golang.org/x/sync/errgroup"
)

func InitDB(fileName string) error {
	var err error
	m.TTTDB, err = bolt.Open(fileName, 0600, nil)
	if err != nil {
		return err
	}

	g := new(errgroup.Group)
	g.Go(func() error { return taskDB.Setup() })
	g.Go(func() error { return sessionDB.Setup() })
	g.Go(func() error { return taskSessionDB.Setup() })

	return g.Wait()
}

func CloseDB() error {
	return m.TTTDB.Close()
}
