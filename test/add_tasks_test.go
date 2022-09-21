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

	addTests([]string{"jogging"})

	tests := []struct {
		args     []string
		expected string
	}{
		{
			[]string{"swimming"}, "swimming",
		},
		{
			[]string{"swimming for 1 hr"}, "swimming for 1 hr",
		},
	}

	for _, test := range tests {
		args := append([]string{"add"}, test.args...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			RootCmd.Execute()
		})

		assert.Contains(t, out, "Added task")
		assert.Contains(t, out, test.expected)
	}
}

func TestBulkAddTaskSuccess(t *testing.T) {
	setup()
	defer teardown()

	addTests([]string{"swimming"})

	tests := []struct {
		args     []string
		expected string
	}{
		{
			[]string{"eating"}, "1. eating",
		},
		{
			[]string{"sleeping", "jogging"}, "1. sleeping\n2. jogging",
		},
	}

	for _, test := range tests {
		args := append([]string{"add", "-b"}, test.args...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			RootCmd.Execute()
		})

		assert.Contains(t, out, "Added task(s)")
		assert.Contains(t, out, test.expected)
	}
}

func TestAddTaskFailure(t *testing.T) {
	setup()
	defer teardown()

	addTests([]string{"swimming", "jogging"})

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

	addTests([]string{"swimming"})

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
