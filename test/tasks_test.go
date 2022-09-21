package test_test

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	. "github.com/xindixu/todo-time-tracker/cmd"
	"github.com/xindixu/todo-time-tracker/db"
	"github.com/zenizh/go-capturer"
)

var testDB string = "todo-time-tracker-test.db"

func setup() {
	db.InitDB(testDB)
}

func teardown() {
	db.CloseDB()
	os.Remove(testDB)
}

func TestAddTaskSuccess(t *testing.T) {
	setup()
	defer teardown()

	for _, test := range []struct {
		args     []string
		expected string
	}{
		{
			[]string{"swimming"}, "swimming",
		},
		{
			[]string{"swimming for 1 hr"}, "swimming for 1 hr",
		},
	} {
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

	for _, test := range []struct {
		args     []string
		expected string
	}{
		{
			[]string{"eating"}, "1. eating",
		},
		{
			[]string{"sleeping", "jogging"}, "1. sleeping\n2. jogging",
		},
	} {
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

	fakeExit := func(int) {
		panic("exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	RootCmd.SetArgs([]string{"add", "swimming"})
	RootCmd.Execute()

	for _, test := range []struct {
		args []string
	}{
		{
			[]string{"swimming"},
		},
	} {
		args := append([]string{"add"}, test.args...)
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

	fakeExit := func(int) {
		panic("exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	RootCmd.SetArgs([]string{"add", "swimming"})
	RootCmd.Execute()

	for _, test := range []struct {
		args []string
	}{
		{
			[]string{"swimming", "jogging"},
		},
		{
			[]string{"jogging", "swimming"},
		},
	} {
		args := append([]string{"add", "-b"}, test.args...)
		RootCmd.SetArgs(args)

		out := capturer.CaptureOutput(func() {
			assert.Panics(t, func() { RootCmd.Execute() })
		})

		assert.Contains(t, out, "Aborted")
		assert.Contains(t, out, "already exists")
	}
}
