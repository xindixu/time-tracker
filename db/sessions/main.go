package taskDB

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	TaskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	m "github.com/xindixu/todo-time-tracker/models"
)

func Setup() error {
	err := m.TTTDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(m.SessionBucketName)
		return err
	})

	return err
}

func update(tx *bolt.Tx, session *m.Session) error {
	sessionBucket := tx.Bucket(m.SessionBucketName)
	taskBucket := tx.Bucket(m.TaskBucketName)
	taskSessionBucket := tx.Bucket(m.TaskSessionBucketName)

	exist, _, err := TaskDB.IsTaskExist(taskBucket, session.Task)
	if !exist {
		return err
	}

	encoded, err := json.Marshal(session)
	if err != nil {
		return err
	}
	err = sessionBucket.Put(m.SessionKey(session.Started), encoded)
	if err != nil {
		return err
	}
	err = taskSessionBucket.Put(m.TaskSessionKey(session.Task, session.Started), []byte{})
	if err != nil {
		return err
	}
	return nil
}

// -----------------------------------

func add(tx *bolt.Tx, started time.Time, task string) (*m.Session, error) {
	session := &m.Session{
		Started: started,
		Ended:   time.Time{},
		Task:    task,
	}

	err := update(tx, session)
	return session, err
}

func AddSession(started time.Time, task string) (*m.Session, error) {
	var session *m.Session
	var err error
	err = m.TTTDB.Update(func(tx *bolt.Tx) error {
		session, err = add(tx, started, task)
		return err
	})
	return session, err
}
