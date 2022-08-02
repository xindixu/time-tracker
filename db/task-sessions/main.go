package taskSessionDB

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	m "github.com/xindixu/todo-time-tracker/models"
)

func Setup() error {
	err := m.TTTDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(m.TaskSessionBucketName)
		return err
	})

	return err
}

func LogAllTaskSessions() ([]m.TaskSession, error) {
	var taskSessions []m.TaskSession

	err := m.TTTDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.TaskSessionBucketName)
		bucket.ForEach(func(k, v []byte) error {
			var taskSession m.TaskSession
			err := json.Unmarshal(v, &taskSession)
			if err != nil {
				return err
			}
			taskSessions = append(taskSessions, taskSession)
			return nil
		})
		return nil
	})
	return taskSessions, err
}
