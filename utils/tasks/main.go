package tasksUtils

import (
	"fmt"
	"os"

	m "github.com/xindixu/todo-time-tracker/models"
)

func Format(task m.Task) string {
	completed := ""
	if !task.Completed.IsZero() {
		completed = "  (done)"
	}

	return fmt.Sprintf("%v%v", task.Title, completed)
}

func Print(task *m.Task, message string) {
	fmt.Printf("%v: %v\n", message, Format(*task))
}

func PrintList(tasks []*m.Task, message string) {
	fmt.Printf("%v:\n", message)

	for i, task := range tasks {
		fmt.Printf("%v. %v\n", i+1, Format(*task))
	}
}

func ActionWithErrorHandling(action func() (*m.Task, error)) *m.Task {
	task, err := action()
	if err != nil {
		fmt.Printf("Aborted command due to:\n%s\n", err)
		os.Exit(1)
	}
	return task
}

func BatchActionWithErrorHandling(action func() ([]*m.Task, error)) []*m.Task {
	tasks, err := action()
	if err != nil {
		fmt.Printf("Aborted command due to:\n%s\n", err)
		os.Exit(1)
	}
	return tasks
}
