package validation

import (
	"errors"
	"taskmanager/schemas"
)

func ValidateTask(task schemas.Task) error {
	if task.Name == "" {
		return errors.New("task name cannot be empty")
	}
	return nil
}
