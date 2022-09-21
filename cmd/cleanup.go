/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	taskDB "github.com/xindixu/todo-time-tracker/db/tasks"
	taskUtil "github.com/xindixu/todo-time-tracker/utils/tasks"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove all completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Removing all completed tasks...\n")

		err := taskDB.CleanupTasks()
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			os.Exit(1)
		}

		tasks, err := taskDB.ListTasks()
		if err != nil {
			return
		}

		if len(tasks) == 0 {
			fmt.Printf("Removed all completed tasks. Nothing left now!\n")
			return
		}

		fmt.Printf("Removed all completed tasks. Here's the remaining tasks:\n")
		i := 0
		for _, task := range tasks {
			if !task.Completed.IsZero() {
				continue
			}
			i += 1
			fmt.Printf("%v. %v\n", i, taskUtil.Format(task))
		}
	},
}

func init() {
	RootCmd.AddCommand(cleanupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
