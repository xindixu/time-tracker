package sessionsUtils

import (
	"fmt"
	"os"

	m "github.com/xindixu/todo-time-tracker/models"
)

func ActionWithErrorHandling(action func() (*m.Session, error)) *m.Session {
	session, err := action()
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err)
		os.Exit(1)
	}
	return session
}
