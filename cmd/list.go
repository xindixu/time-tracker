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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List out all of added tasks",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing all task...\n")

		tasks, err := taskDB.ListTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Printf("You haven't added any task yet\n")
			return
		}

		fmt.Printf("Here's a list of all your incomplete tasks\n")
		for i, task := range tasks {
			fmt.Printf("%v. %v\n", i+1, taskUtil.Format(task))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
