package sessionsUtils

import (
	"fmt"
	"os"
	"time"

	m "github.com/xindixu/todo-time-tracker/models"
)

func Format(session m.Session) string {
	end := "???"
	if !session.Ended.IsZero() {
		end = session.Ended.Format(time.RubyDate)
	}
	return fmt.Sprintf("%v - %v:  %v", session.Started.Format(time.RubyDate), end, session.Task)
}

func ActionWithErrorHandling(action func() (*m.Session, error)) *m.Session {
	session, err := action()
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err)
		os.Exit(1)
	}
	return session
}
