package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	ID        int
	Created   time.Time
	Completed time.Time
	Title     string
}

var taskBucket = []byte("tasks")
var db *bolt.DB

func AddTask(title string) error {
	task := &Task{
		Created:   time.Now(),
		Completed: time.Time{},
		Title:     title,
	}

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		id64, _ := bucket.NextSequence()
		task.ID = int(id64)

		encoded, err := json.Marshal(task)
		if err != nil {
			return err
		}
		err = bucket.Put(itob(task.ID), encoded)
		return err
	})

	return err
}

func ListTasks() error {
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
	return err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func setup() error {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

	return err
}

func InitDB() error {
	var err error
	db, err = bolt.Open("todo-time-tracker.db", 0600, nil)
	if err != nil {
		return err
	}

	return setup()
}
