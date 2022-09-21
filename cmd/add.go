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

func addBatch(args []string) {
	args = helperUtils.Dedup(args)
	tasks := taskUtil.BatchActionWithErrorHandling(func() ([]*m.Task, error) { return taskDB.BatchAddTasks(args) })

	taskUtil.PrintList(tasks, "Added task(s)")
}

func addOne(args []string) {
	title := strings.Join(args, " ")

	task := taskUtil.ActionWithErrorHandling(func() (*m.Task, error) { return taskDB.AddTask(title) })

	taskUtil.Print(task, "Added task")
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task or a list of tasks",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding tasks...\n")
	},
	Run: func(cmd *cobra.Command, args []string) {
		batch, _ := cmd.Flags().GetBool("batch")
		if batch {
			addBatch(args)
		} else {
			addOne(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().BoolP("batch", "b", false, "Add multiple tasks")
}
