package tasksUtils_test

import (
	"testing"
	"time"

	m "github.com/xindixu/todo-time-tracker/models"
	tasksUtils "github.com/xindixu/todo-time-tracker/utils/tasks"
)

func TestFormat1(t *testing.T) {
	created, _ := time.Parse(time.RFC3339, "2022-09-20T15:00:00Z")
	completed, _ := time.Parse(time.RFC3339, "2022-09-20T16:00:00Z")

	task := m.Task{
		Created:   created,
		Completed: completed,
		Title:     "learn go",
	}

	expected := "learn go ✓"
	s := tasksUtils.Format(task)
	if s != expected {
		t.Errorf("\nExpected: %s\nGot: %s", expected, s)
	}
}

func TestFormat2(t *testing.T) {
	created, _ := time.Parse(time.RFC3339, "2022-09-20T15:00:00Z")

	task := m.Task{
		Created:   created,
		Completed: time.Time{},
		Title:     "learn go",
	}

	expected := "learn go"
	s := tasksUtils.Format(task)
	if s != expected {
		t.Errorf("\nExpected: %s\nGot: %s", expected, s)
	}
}
