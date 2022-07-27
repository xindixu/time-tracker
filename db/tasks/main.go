package taskDB

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"time"

	"github.com/boltdb/bolt"
	"github.com/xindixu/todo-time-tracker/models"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

func AddTask(title string) (*models.Task, error) {
	task := &models.Task{
		Created:   time.Now(),
		Completed: time.Time{},
		Title:     title,
	}

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		key, err := getTaskKeyByTitle(bucket, title)
		if err != nil {
			return err
		}
		if key != nil {
			return fmt.Errorf("task \"%s\" already exists", title)
		}

		id64, _ := bucket.NextSequence()
		task.ID = int(id64)

		encoded, err := json.Marshal(task)
		if err != nil {
			return err
		}
		err = bucket.Put(itob(task.ID), encoded)
		return err
	})

	if err != nil {
		return nil, err
	}
	return task, err
}

func CompleteTask(title string) (*models.Task, error) {
	var task models.Task
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		key, err := getTaskKeyByTitle(bucket, title)
		if err != nil {
			return err
		}
		if key == nil {
			return fmt.Errorf("task \"%s\" not found", title)
		}

		v := bucket.Get(key)

		err = json.Unmarshal(v, &task)
		if err != nil {
			return err
		}

		task.Completed = time.Now()
		encoded, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return bucket.Put(key, encoded)

	})

	return &task, err
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

func getTaskKeyByTitle(bucket *bolt.Bucket, title string) ([]byte, error) {
	var key []byte

	cursor := bucket.Cursor()

	for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
		var task models.Task
		err := json.Unmarshal(v, &task)
		if err != nil {
			return nil, err
		}

		if task.Title == title {
			key = k
			return key, nil
		}
	}

	return nil, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func Setup(baseDb *bolt.DB) error {
	db = baseDb
	err := baseDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

	return err
}
