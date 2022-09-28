package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/xindixu/todo-time-tracker/cmd"
	"github.com/zenizh/go-capturer"
)

func TestCompleteTaskSuccess(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)

	tests := [][]string{
		{tasks[0]},
		{tasks[1]},
	}

	// Complete tasks and test output
	for _, test := range tests {
		args := append([]string{"done"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			RootCmd.Execute()
		})

		assert.Contains(t, out, "Completed task")
		for _, task := range test {
			assert.Contains(t, out, task)
		}
	}

	// Test if tasks are completed in db
	allTasks := getTasks()
	for _, test := range tests {
		for _, task := range test {
			assert.Contains(t, allTasks, fmt.Sprintf("%s ✓", task))
		}
	}
}

func TestBulkCompleteTaskSuccess(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)

	tests := [][]string{
		tasks[0:1],
		tasks[1:3],
	}

	// Complete tasks and test output
	for _, test := range tests {
		args := append([]string{"done", "-b"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			RootCmd.Execute()
		})

		assert.Contains(t, out, "Completed task(s)")
		for _, task := range test {
			assert.Contains(t, out, task)
		}
	}

	// Test if tasks are completed in db
	allTasks := getTasks()
	for _, test := range tests {
		for _, task := range test {
			assert.Contains(t, allTasks, fmt.Sprintf("%s ✓", task))
		}
	}
}

func TestCompleteTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)

	tests := [][]string{
		{"studying"},
		{"reading for 1 hr"},
	}

	for _, test := range tests {
		args := append([]string{"done"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			assert.Panics(t, func() { RootCmd.Execute() })
		})

		assert.Contains(t, out, "Aborted")
		assert.Contains(t, out, "not found")
	}
}

func TestBulkCompleteTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)

	tests := [][]string{
		{"studying", "reading"},
		{"reading for 1 hr", "coding"},
	}

	for _, test := range tests {
		args := append([]string{"done", "-b"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			assert.Panics(t, func() { RootCmd.Execute() })
		})

		assert.Contains(t, out, "Aborted")
		assert.Contains(t, out, "not found")
	}
}

func TestRepeatedCompleteTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)
	test := tasks[0:1]
	completeTasks(test)

	out := capturer.CaptureOutput(func() {
		assert.Panics(t, func() {
			completeTasks(test)
		})
	})

	assert.Contains(t, out, "Aborted")
	assert.Contains(t, out, "already marked as completed")
}
