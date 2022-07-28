/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	sessionDB "github.com/xindixu/todo-time-tracker/db/sessions"
	m "github.com/xindixu/todo-time-tracker/models"
	sessionUtil "github.com/xindixu/todo-time-tracker/utils/sessions"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the tasking that is currently being worked on",
	Run: func(cmd *cobra.Command, args []string) {
		session := sessionUtil.ActionWithErrorHandling(func() (*m.Session, error) {
			return sessionDB.GetActiveSession()
		})

		fmt.Printf("Current task %s started at %s.\n", session.Task, session.Started.Format(time.RubyDate))
		fmt.Printf("You have been working on it for %s.\n", time.Since(session.Started).Round(time.Second))
		fmt.Printf("Keep up the good work!\n")
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
