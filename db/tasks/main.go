package taskDB

import (
	"encoding/json"
	"fmt"

	"time"

	"github.com/boltdb/bolt"
	h "github.com/xindixu/todo-time-tracker/db/helper"
	sessionDB "github.com/xindixu/todo-time-tracker/db/sessions"
	taskSessionDB "github.com/xindixu/todo-time-tracker/db/task-sessions"
	m "github.com/xindixu/todo-time-tracker/models"
	"golang.org/x/sync/errgroup"
)

func Setup() error {
	err := m.TTTDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(m.TaskBucketName)
		return err
	})

	return err
}

func update(bucket *bolt.Bucket, task *m.Task) error {
	encoded, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return bucket.Put(m.TaskKey(task.Title), encoded)
}

// -----------------------------------

func add(bucket *bolt.Bucket, title string) (*m.Task, error) {
	exist, _, err := h.IsTaskExist(bucket, title)
	if exist {
		return nil, err
	}

	task := &m.Task{
		Created:   time.Now(),
		Completed: time.Time{},
		Title:     title,
	}

	err = update(bucket, task)
	return task, err
}

func complete(bucket *bolt.Bucket, title string) (*m.Task, error) {
	exist, v, err := h.IsTaskExist(bucket, title)
	if !exist {
		return nil, err
	}

	var task m.Task
	err = json.Unmarshal(v, &task)
	if err != nil {
		return nil, err
	}
	if !task.Completed.IsZero() {
		return nil, fmt.Errorf("task \"%s\" is already marked as complete", title)
	}
	task.Completed = time.Now()
	err = update(bucket, &task)
	return &task, err
}

func delete(tx *bolt.Tx, title string) (*m.Task, error) {
	taskBucket := tx.Bucket(m.TaskBucketName)
	sessionBucket := tx.Bucket(m.SessionBucketName)
	taskSessionBucket := tx.Bucket(m.TaskSessionBucketName)

	exist, v, err := h.IsTaskExist(taskBucket, title)
	if !exist {
		return nil, err
	}

	var task m.Task
	err = json.Unmarshal(v, &task)
	if err != nil {
		return nil, err
	}

	err = taskBucket.Delete(m.TaskKey(title))

	if err != nil {
		return nil, err
	}
	taskSessionKeys, sessionKeys := taskSessionDB.GetTaskSessionKeysByTask(taskSessionBucket, title)

	err = taskSessionDB.BatchDeleteTaskSessions(taskSessionBucket, taskSessionKeys)
	if err != nil {
		return nil, err
	}
	err = sessionDB.BatchDeleteSessions(sessionBucket, sessionKeys)
	if err != nil {
		return nil, err
	}

	return &task, err
}

// -----------------------------------

func AddTask(title string) (*m.Task, error) {
	var task *m.Task
	var err error
	err = m.TTTDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.TaskBucketName)
		task, err = add(bucket, title)
		return err
	})
	return task, err
}

func CompleteTask(title string) (*m.Task, error) {
	var task *m.Task
	var err error
	err = m.TTTDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.TaskBucketName)
		task, err = complete(bucket, title)
		return err
	})
	return task, err
}

func DeleteTask(title string) (*m.Task, error) {
	var task *m.Task
	var err error
	err = m.TTTDB.Update(func(tx *bolt.Tx) error {
		task, err = delete(tx, title)
		return err
	})
	return task, err
}

// -----------------------------------

func BatchAddTasks(titles []string) ([]*m.Task, error) {
	tasks := make([]*m.Task, len(titles))

	err := m.TTTDB.Batch(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.TaskBucketName)

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

func BatchCompleteTasks(titles []string) ([]*m.Task, error) {
	tasks := make([]*m.Task, len(titles))
	err := m.TTTDB.Batch(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.TaskBucketName)

		g := new(errgroup.Group)

		for i, title := range titles {
			func(i int, title string) {
				g.Go(func() error {
					task, err := complete(bucket, title)
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

func BatchDeleteTasks(titles []string) ([]*m.Task, error) {
	tasks := make([]*m.Task, len(titles))
	err := m.TTTDB.Batch(func(tx *bolt.Tx) error {

		g := new(errgroup.Group)

		for i, title := range titles {
			func(i int, title string) {
				g.Go(func() error {
					task, err := delete(tx, title)
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

func CleanupTasks() error {
	err := m.TTTDB.Batch(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.TaskBucketName)

		err := bucket.ForEach(func(k, v []byte) error {
			var task m.Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			if !task.Completed.IsZero() {
				err = bucket.Delete(m.TaskKey(task.Title))
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

func ListTasks() ([]m.Task, error) {
	var tasks []m.Task

	err := m.TTTDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.TaskBucketName)
		bucket.ForEach(func(k, v []byte) error {
			var task m.Task
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
