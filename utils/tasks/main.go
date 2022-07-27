package tasksUtils

import (
	"fmt"

	"github.com/xindixu/todo-time-tracker/models"
)

func Format(task models.Task) string {
	completed := ""
	if !task.Completed.IsZero() {
		completed = "  (done)"
	}

	return fmt.Sprintf("%v %v", task.Title, completed)
}
