package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	sessionDB "github.com/xindixu/todo-time-tracker/db/sessions"
	m "github.com/xindixu/todo-time-tracker/models"
	sessionUtil "github.com/xindixu/todo-time-tracker/utils/sessions"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the timer for a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		fmt.Printf("Starting a new session for task: %s...\n", task)

		now := time.Now()
		session := sessionUtil.ActionWithErrorHandling(func() (*m.Session, error) {
			return sessionDB.StartSession(now, task)
		})

		fmt.Printf("Started task %s at %s. Have fun!\n", session.Task, session.Started.Format(time.RubyDate))
	},
}

func init() {
	RootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
