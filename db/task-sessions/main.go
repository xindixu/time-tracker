package taskDB

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var taskSessionBucket = []byte("taskSessions")
var db *bolt.DB

func key(title string, started time.Time) []byte {
	return []byte(fmt.Sprintf("%v:%v", title, started.Format(time.RFC3339)))
}

func Setup(baseDb *bolt.DB) error {
	db = baseDb
	err := baseDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskSessionBucket)
		return err
	})

	return err
}
