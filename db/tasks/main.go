package taskDB

import (
	"encoding/json"
	"fmt"

	"time"

	"github.com/boltdb/bolt"
	"github.com/xindixu/todo-time-tracker/models"
	"golang.org/x/sync/errgroup"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

func update(bucket *bolt.Bucket, task *models.Task) error {
	encoded, err := json.Marshal(task)
	if err != nil {
		return err
	}
	err = bucket.Put([]byte(task.Title), encoded)
	if err != nil {
		return err
	}
	return nil
}

// -----------------------------------

func add(bucket *bolt.Bucket, title string) (*models.Task, error) {
	v := bucket.Get(key(title))
	if v != nil {
		return nil, fmt.Errorf("task \"%s\" already exists", title)
	}

	task := &models.Task{
		Created:   time.Now(),
		Completed: time.Time{},
		Title:     title,
	}

	err := update(bucket, task)
	return task, err
}

func AddTask(title string) (*models.Task, error) {
	var task *models.Task
	var err error
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		task, err = add(bucket, title)
		return err
	})
	return task, err
}

func BatchAddTasks(titles []string) ([]*models.Task, error) {
	tasks := make([]*models.Task, len(titles))

	err := db.Batch(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		g := new(errgroup.Group)

		for i, title := range titles {
			func(i int, title string) {
				g.Go(func() error {
					task, err := add(bucket, title)
					tasks[i] = task
					return err
				})
			}(i, title)
		}
		if err := g.Wait(); err != nil {
			return err
		}

		return nil
	})

	return tasks, err
}

func DeleteTask(title string) (*models.Task, error) {
	var task models.Task
	err := db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(taskBucket)
		v := bucket.Get(key(title))
		if v == nil {
			return fmt.Errorf("task \"%s\" not found", title)
		}

		err := json.Unmarshal(v, &task)
		if err != nil {
			return err
		}

		return bucket.Delete(key(title))
	})

	return &task, err
}

func CleanupTasks() error {
	err := db.Batch(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		err := bucket.ForEach(func(k, v []byte) error {
			var task models.Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			if !task.Completed.IsZero() {
				err = bucket.Delete(key(task.Title))
				if err != nil {
					return err
				}
			}
			return nil
		})

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

func key(s string) []byte {
	return []byte(s)
}

func Setup(baseDb *bolt.DB) error {
	db = baseDb
	err := baseDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

	return err
}
