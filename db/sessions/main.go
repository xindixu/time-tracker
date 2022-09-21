package sessionDB

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	h "github.com/xindixu/todo-time-tracker/db/helper"
	m "github.com/xindixu/todo-time-tracker/models"
)

func Setup() error {
	err := m.TTTDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(m.SessionBucketName)
		return err
	})

	return err
}

func update(bucket *bolt.Bucket, session *m.Session) error {
	encoded, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return bucket.Put(m.SessionKey(session.Started), encoded)
}

// -----------------------------------

func start(tx *bolt.Tx, started time.Time, task string) (*m.Session, error) {
	taskBucket := tx.Bucket(m.TaskBucketName)
	exist, _, err := h.IsTaskExist(taskBucket, task)
	if !exist {
		return nil, err
	}

	session := &m.Session{
		Started: started,
		Ended:   time.Time{},
		Task:    task,
	}

	err = update(tx.Bucket(m.SessionBucketName), session)
	if err != nil {
		return session, err
	}

	taskSessionBucket := tx.Bucket(m.TaskSessionBucketName)
	err = taskSessionBucket.Put(m.TaskSessionKey(session.Task, session.Started), []byte{})
	if err != nil {
		return session, err
	}

	return session, err
}

func HasActiveSession(bucket *bolt.Bucket) (bool, []byte, error) {
	_, v := bucket.Cursor().Last()
	if v == nil {
		return false, nil, fmt.Errorf("no session found")
	}
	var session m.Session
	err := json.Unmarshal(v, &session)
	if err != nil {
		return false, nil, fmt.Errorf("can't unmarshal session")
	}
	if session.Ended.IsZero() {
		return true, v, fmt.Errorf("already has an active session")
	}
	return false, v, fmt.Errorf("no active session")
}

func BatchDeleteSessions(bucket *bolt.Bucket, sessionKeys [][]byte) error {
	for _, key := range sessionKeys {
		return bucket.Delete(key)
	}
	return nil
}

// -----------------------------------

func StartSession(started time.Time, task string) (*m.Session, error) {
	var session *m.Session
	err := m.TTTDB.Update(func(tx *bolt.Tx) error {
		active, _, err := HasActiveSession(tx.Bucket(m.SessionBucketName))
		if active {
			return err
		}

		session, err = start(tx, started, task)
		return err
	})
	return session, err
}

func GetActiveSession() (*m.Session, error) {
	var session m.Session

	err := m.TTTDB.View(func(tx *bolt.Tx) error {
		active, v, err := HasActiveSession(tx.Bucket(m.SessionBucketName))
		if !active {
			return err
		}

		err = json.Unmarshal(v, &session)
		if err != nil {
			return err
		}
		return err
	})
	return &session, err
}

func EndSession(ended time.Time) (*m.Session, error) {
	session, err := GetActiveSession()
	if err != nil {
		return session, err
	}

	err = m.TTTDB.Update(func(tx *bolt.Tx) error {
		session.Ended = ended
		bucket := tx.Bucket(m.SessionBucketName)
		return update(bucket, session)
	})
	return session, err
}

func ListSessions() ([]m.Session, error) {
	var sessions []m.Session

	err := m.TTTDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(m.SessionBucketName)
		bucket.ForEach(func(k, v []byte) error {
			var session m.Session
			err := json.Unmarshal(v, &session)
			if err != nil {
				return err
			}
			sessions = append(sessions, session)
			return nil
		})
		return nil
	})
	return sessions, err
}
