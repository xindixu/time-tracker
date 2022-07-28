package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	m "github.com/xindixu/todo-time-tracker/models"
	taskUtil "github.com/xindixu/todo-time-tracker/utils/tasks"
)

func doneBatch(args []string) {
	fmt.Printf("Completing task(s): %s...\n", strings.Join(args, ", "))

	tasks := taskUtil.BatchActionWithErrorHandling(func() ([]*m.Task, error) { return taskDB.BatchCompleteTasks(args) })

	fmt.Printf("Completed task(s):\n")
	for i, task := range tasks {
		fmt.Printf("%v. %v\n", i+1, taskUtil.Format(*task))
	}
}

func doneOne(args []string) {
	title := strings.Join(args, " ")
	fmt.Printf("Completing task: %s...\n", title)

	task := taskUtil.ActionWithErrorHandling(func() (*m.Task, error) { return taskDB.CompleteTask(title) })
	fmt.Printf("Completed task: %v\n", taskUtil.Format(*task))
}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task or a list of tasks as completed",
	Args:  cobra.MinimumNArgs(1),
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
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	doneCmd.Flags().BoolP("batch", "b", false, "Mark multiple tasks as done")
}
