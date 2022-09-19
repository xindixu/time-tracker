package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	m "github.com/xindixu/todo-time-tracker/models"
	taskUtil "github.com/xindixu/todo-time-tracker/utils/tasks"
)

func deleteBatch(args []string) {
	tasks := taskUtil.BatchActionWithErrorHandling(func() ([]*m.Task, error) { return taskDB.BatchDeleteTasks(args) })

	taskUtil.PrintList(tasks, "Deleted task(s)")
}

func deleteOne(args []string) {
	title := strings.Join(args, " ")

	task := taskUtil.ActionWithErrorHandling(func() (*m.Task, error) { return taskDB.DeleteTask(title) })

	taskUtil.Print(task, "Deleted task")
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task or a list of tasks",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deleting tasks...\n")
	},
	Run: func(cmd *cobra.Command, args []string) {
		batch, _ := cmd.Flags().GetBool("batch")
		if batch {
			deleteBatch(args)
		} else {
			deleteOne(args)
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().BoolP("batch", "b", false, "Delete multiple tasks")
}
