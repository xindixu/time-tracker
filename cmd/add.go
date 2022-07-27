package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	"github.com/xindixu/todo-time-tracker/models"
	taskUtil "github.com/xindixu/todo-time-tracker/utils/tasks"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task or a list of tasks",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding task(s): %s...\n", strings.Join(args, ", "))
		var tasks []models.Task

		for _, v := range args {
			task, err := taskDB.AddTask(v)
			if err != nil {
				fmt.Printf("Something went wrong: %s\n", err)
				os.Exit(1)
			}
			if task != nil {
				tasks = append(tasks, *task)
			}
		}

		fmt.Printf("Added task(s):\n")
		for i, task := range tasks {
			fmt.Printf("%v. %v\n", i+1, taskUtil.Format(task))
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
}
