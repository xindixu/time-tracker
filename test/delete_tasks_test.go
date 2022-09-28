package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/xindixu/todo-time-tracker/cmd"
	"github.com/zenizh/go-capturer"
)

func TestDeleteTaskSuccess(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)

	tests := [][]string{
		{tasks[0]},
		{tasks[1]},
	}

	// Delete tasks and test output
	for _, test := range tests {
		args := append([]string{"delete"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			RootCmd.Execute()
		})

		assert.Contains(t, out, "Deleted task")
		for _, task := range test {
			assert.Contains(t, out, task)
		}
	}

	// Test if tasks are deleted in db
	allTasks := getTasks()
	for _, test := range tests {
		for _, task := range test {
			assert.NotContains(t, allTasks, task)
		}
	}
}

func TestBulkDeleteTaskSuccess(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)

	tests := [][]string{
		tasks[0:1],
		tasks[1:3],
	}

	// Delete tasks and test output
	for _, test := range tests {
		args := append([]string{"delete", "-b"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			RootCmd.Execute()
		})

		assert.Contains(t, out, "Deleted task(s)")
		for _, task := range test {
			assert.Contains(t, out, task)
		}
	}

	// Test if tasks are deleted in db
	allTasks := getTasks()
	for _, test := range tests {
		for _, task := range test {
			assert.NotContains(t, allTasks, task)
		}
	}
}

func TestDeleteTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)

	tests := [][]string{
		{"studying"},
		{"reading for 1 hr"},
	}

	for _, test := range tests {
		args := append([]string{"delete"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			assert.Panics(t, func() { RootCmd.Execute() })
		})

		assert.Contains(t, out, "Aborted")
		assert.Contains(t, out, "not found")
	}
}

func TestBulkDeleteTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)

	tests := [][]string{
		{"studying", "reading"},
		{"reading for 1 hr", "coding"},
	}

	for _, test := range tests {
		args := append([]string{"delete", "-b"}, test...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			assert.Panics(t, func() { RootCmd.Execute() })
		})

		assert.Contains(t, out, "Aborted")
		assert.Contains(t, out, "not found")
	}
}

func TestRepeatedDeleteTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTasks(tasks)
	test := tasks[0:1]
	deleteTasks(test)

	out := capturer.CaptureOutput(func() {
		assert.Panics(t, func() {
			deleteTasks(test)
		})
	})

	assert.Contains(t, out, "Aborted")
	assert.Contains(t, out, "not found")
}
