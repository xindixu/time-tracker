package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	m "github.com/xindixu/todo-time-tracker/models"
	helperUtils "github.com/xindixu/todo-time-tracker/utils/helper"
	taskUtil "github.com/xindixu/todo-time-tracker/utils/tasks"
)

func doneBatch(args []string) {
	args = helperUtils.Dedup(args)
	tasks := taskUtil.BatchActionWithErrorHandling(func() ([]*m.Task, error) { return taskDB.BatchCompleteTasks(args) })

	taskUtil.PrintList(tasks, "Completed task(s)")
}

func doneOne(args []string) {
	title := strings.Join(args, " ")

	task := taskUtil.ActionWithErrorHandling(func() (*m.Task, error) { return taskDB.CompleteTask(title) })

	taskUtil.Print(task, "Completed task")
}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task or a list of tasks as completed",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Completing tasks...\n")
	},
	Run: func(cmd *cobra.Command, args []string) {
		batch, _ := cmd.Flags().GetBool("batch")
		if batch {
			doneBatch(args)
		} else {
			doneOne(args)
		}

	},
}

func init() {
	RootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	doneCmd.Flags().BoolP("batch", "b", false, "Mark multiple tasks as done")
}
