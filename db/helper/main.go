package dbHelper

import (
	"fmt"

	"github.com/boltdb/bolt"
	m "github.com/xindixu/todo-time-tracker/models"
)

func IsTaskExist(taskBucket *bolt.Bucket, title string) (bool, []byte, error) {
	v := taskBucket.Get(m.TaskKey(title))
	if v != nil {
		return true, v, fmt.Errorf("task \"%s\" already exists", title)
	}
	return false, v, fmt.Errorf("task \"%s\" not found", title)
}
