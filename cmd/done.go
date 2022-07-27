/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	"github.com/xindixu/todo-time-tracker/models"
	taskUtil "github.com/xindixu/todo-time-tracker/utils/tasks"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task or a list of tasks as completed",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []models.Task
		for _, v := range args {
			task, err := taskDB.CompleteTask(v)
			if err != nil {
				fmt.Printf("Something went wrong: %s\n", err)
				os.Exit(1)
			}
			tasks = append(tasks, task)
		}

		fmt.Printf("Completed task(s):\n")
		for i, task := range tasks {
			fmt.Printf("%v. %v\n", i+1, taskUtil.Format(task))
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
}
