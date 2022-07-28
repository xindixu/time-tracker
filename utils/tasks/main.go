package tasksUtils

import (
	"fmt"
	"os"

	"github.com/xindixu/todo-time-tracker/models"
)

func Format(task models.Task) string {
	completed := ""
	if !task.Completed.IsZero() {
		completed = "  (done)"
	}

	return fmt.Sprintf("%v%v", task.Title, completed)
}

func ActionWithErrorHandling(action func() (*models.Task, error)) *models.Task {
	task, err := action()
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err)
		os.Exit(1)
	}
	return task
}

func BatchActionWithErrorHandling(action func() ([]*models.Task, error)) []*models.Task {
	tasks, err := action()
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err)
		os.Exit(1)
	}
	return tasks
}
