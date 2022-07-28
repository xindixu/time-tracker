package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	"github.com/xindixu/todo-time-tracker/models"
	taskUtil "github.com/xindixu/todo-time-tracker/utils/tasks"
)

func addBatch(args []string) {
	fmt.Printf("Adding task(s): %s...\n", strings.Join(args, ", "))

	tasks := taskUtil.BatchActionWithErrorHandling(func() ([]*models.Task, error) { return taskDB.BatchAddTasks(args) })

	fmt.Printf("Added task(s):\n")
	for i, task := range tasks {
		fmt.Printf("%v. %v\n", i+1, taskUtil.Format(*task))
	}
}

func addOne(args []string) {
	title := strings.Join(args, " ")
	fmt.Printf("Adding task: %s...\n", title)

	task := taskUtil.ActionWithErrorHandling(func() (*models.Task, error) { return taskDB.AddTask(title) })
	fmt.Printf("Added task: %v\n", taskUtil.Format(*task))
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task or a list of tasks",
	Args:  cobra.MinimumNArgs(1),
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
