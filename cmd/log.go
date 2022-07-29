/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	sessionDB "github.com/xindixu/todo-time-tracker/db/sessions"
	sessionsUtils "github.com/xindixu/todo-time-tracker/utils/sessions"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show the log of sessions spent on tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Logging all sessions...\n")

		sessions, err := sessionDB.ListSessions()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		if len(sessions) == 0 {
			fmt.Printf("You haven't worked on any task yet\n")
			return
		}

		for _, session := range sessions {
			fmt.Printf("%v\n", sessionsUtils.Format(session))

		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
