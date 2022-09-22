package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/xindixu/todo-time-tracker/cmd"
	"github.com/zenizh/go-capturer"
)

func TestAddTaskSuccess(t *testing.T) {
	setup()
	defer teardown()

	addTasks([]string{"jogging"})

	tests := [][]string{
		{"swimming"},
		{"swimming for 1 hr"},
	}

	// Add tasks and test output
	for _, test := range tests {
		args := append([]string{"add"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			RootCmd.Execute()
		})

		assert.Contains(t, out, "Added task")
		for _, task := range test {
			assert.Contains(t, out, task)
		}
	}

	// Test if tasks are added to db
	allTasks := getTasks()
	for _, test := range tests {
		for _, task := range test {
			assert.Contains(t, allTasks, task)
		}
	}
}

func TestBulkAddTaskSuccess(t *testing.T) {
	setup()
	defer teardown()

	addTasks([]string{"swimming"})

	tests := [][]string{
		{"eating"},
		{"sleeping", "jogging"},
	}

	// Add tasks and test output
	for _, test := range tests {
		args := append([]string{"add", "-b"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			RootCmd.Execute()
		})

		assert.Contains(t, out, "Added task(s)")

		for _, task := range test {
			assert.Contains(t, out, task)
		}
	}

	// Test if tasks are added to db
	allTasks := getTasks()
	for _, test := range tests {
		for _, task := range test {
			assert.Contains(t, allTasks, task)
		}
	}
}

func TestAddTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTasks([]string{"swimming", "jogging"})

	tests := [][]string{
		{"swimming"},
		{"jogging"},
	}

	for _, test := range tests {
		args := append([]string{"add"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			assert.Panics(t, func() { RootCmd.Execute() })
		})

		assert.Contains(t, out, "Aborted")
		assert.Contains(t, out, "already exists")
	}
}

func TestBulkAddTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTasks([]string{"swimming"})

	tests := [][]string{
		{"swimming", "jogging"},
		{"jogging", "swimming"},
	}

	for _, test := range tests {
		args := append([]string{"add", "-b"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			assert.Panics(t, func() { RootCmd.Execute() })
		})

		assert.Contains(t, out, "Aborted")
		assert.Contains(t, out, "already exists")
	}
}
