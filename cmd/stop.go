package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	sessionDB "github.com/xindixu/todo-time-tracker/db/sessions"
	m "github.com/xindixu/todo-time-tracker/models"
	sessionUtil "github.com/xindixu/todo-time-tracker/utils/sessions"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the timer for the current task",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		session := sessionUtil.ActionWithErrorHandling(func() (*m.Session, error) {
			return sessionDB.EndSession(now)
		})

		fmt.Printf("Stopped task %s at %s.\n", session.Task, session.Ended.Format(time.RubyDate))
		fmt.Printf("Time spent: %s.\n", session.Ended.Sub(session.Started).Round(time.Second))
		fmt.Printf("Great work! Now go take some rest.\n")
	},
}

func init() {
	RootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
