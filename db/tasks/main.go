package taskDB

import (
	"encoding/binary"
	"encoding/json"

	"time"

	"github.com/boltdb/bolt"
	"github.com/xindixu/todo-time-tracker/models"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

func AddTask(title string) error {
	task := &models.Task{
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

func ListTasks() ([]models.Task, error) {
	var tasks []models.Task

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		bucket.ForEach(func(k, v []byte) error {
			var task models.Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			tasks = append(tasks, task)
			return nil
		})
		return nil
	})
	return tasks, err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func Setup(baseDb *bolt.DB) error {
	db = baseDb
	err := baseDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

	return err
}
