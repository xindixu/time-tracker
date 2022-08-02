package taskSessionDB

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	m "github.com/xindixu/todo-time-tracker/models"
	"golang.org/x/sync/errgroup"
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

func GetTaskSessionKeysByTask(bucket *bolt.Bucket, task string) ([][]byte, [][]byte) {
	var sessionKeys [][]byte
	var taskSessionKeys [][]byte

	cursor := bucket.Cursor()
	prefix := m.TaskSessionKey(task, time.Time{})
	for k, _ := cursor.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, _ = cursor.Next() {
		taskSessionKeys = append(taskSessionKeys, k)
		sessionKeys = append(sessionKeys, m.SessionKeyFromTaskSessionKey(k))
	}
	return taskSessionKeys, sessionKeys
}

func BatchDeleteTaskSessions(bucket *bolt.Bucket, taskSessionKeys [][]byte) error {
	g := new(errgroup.Group)

	for _, key := range taskSessionKeys {
		func(key []byte) {
			g.Go(func() error {
				return bucket.Delete(key)
			})
		}(key)
	}

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
