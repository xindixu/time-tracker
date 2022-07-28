package taskDB

import (
	"time"

	"github.com/boltdb/bolt"
)

var sessionBucket = []byte("sessions")
var db *bolt.DB

func key(started time.Time) []byte {
	return []byte(started.Format(time.RFC3339))
}

func Setup(baseDb *bolt.DB) error {
	db = baseDb
	err := baseDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(sessionBucket)
		return err
	})

	return err
}
