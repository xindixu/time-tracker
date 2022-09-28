package test

import (
	"os"

	"bou.ke/monkey"
	. "github.com/xindixu/todo-time-tracker/cmd"
	"github.com/xindixu/todo-time-tracker/db"
	"github.com/zenizh/go-capturer"
)

var testDB string = "todo-time-tracker-test.db"
var fakeExit = func(int) {
	panic("exit called")
}

var patch *monkey.PatchGuard

func setup() {
	db.InitDB(testDB)

	patch = monkey.Patch(os.Exit, fakeExit)
}

func teardown() {
	db.CloseDB()
	os.Remove(testDB)

	patch.Unpatch()
}

func addTasks(tasks []string) {
	RootCmd.SetArgs(append([]string{"add", "-b"}, tasks...))
	RootCmd.Execute()
}

func getTasks() string {
	RootCmd.SetArgs([]string{"list", "-a"})
	out := capturer.CaptureOutput(func() {
		RootCmd.Execute()
	})
	return out
}

func completeTasks(tasks []string) {
	RootCmd.SetArgs(append([]string{"done", "-b"}, tasks...))
	RootCmd.Execute()
}

func deleteTasks(tasks []string) {
	RootCmd.SetArgs(append([]string{"delete", "-b"}, tasks...))
	RootCmd.Execute()
}

var tasks = []string{"swimming", "swimming for 1 hr", "sleeping", "jogging"}
