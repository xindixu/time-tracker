package models

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Created   time.Time
	Completed time.Time
	Title     string // PK
}

type Session struct {
	Started time.Time // PK
	Ended   time.Time
	Task    string // FK(Task.Title)
}

type TaskSession struct {
	Task    string    // FK(Task.Title)
	Session time.Time // FK(Session.Started)
}

var TTTDB *bolt.DB

var TaskBucketName = []byte("tasks")
var SessionBucketName = []byte("sessions")
var TaskSessionBucketName = []byte("TaskSessions")

func TaskKey(title string) []byte {
	return []byte(title)
}
func SessionKey(started time.Time) []byte {
	return []byte(started.Format(time.RFC3339))
}
func TaskSessionKey(title string, started time.Time) []byte {
	if started.IsZero() {
		return []byte(fmt.Sprintf("%v:", title))
	}
	return []byte(fmt.Sprintf("%v:%v", title, started.Format(time.RFC3339)))
}

func SessionKeyFromTaskSessionKey(key []byte) []byte {
	return []byte(string(key)[len(string(key))-len(time.RFC3339):])
}

func BtoTime(b []byte) time.Time {
	str := string(b)
	t, _ := time.Parse(time.RFC3339, str)
	return t
}
