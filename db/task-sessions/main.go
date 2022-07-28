package taskDB

import (
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
