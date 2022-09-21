package test

import (
	"os"

	"bou.ke/monkey"
	"github.com/xindixu/todo-time-tracker/cmd"
	"github.com/xindixu/todo-time-tracker/db"
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

func addTests(tasks []string) {
	cmd.RootCmd.SetArgs(append([]string{"add", "-b"}, tasks...))
	cmd.RootCmd.Execute()
}
