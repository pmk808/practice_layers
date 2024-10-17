package validation

import (
    "errors"
    "taskmanager/schemas"
)

type TaskValidator struct {}

func NewTaskValidator() *TaskValidator {
    return &TaskValidator{}
}

func (v *TaskValidator) ValidateTask(task schemas.Task) error {
    if task.Name == "" {
        return errors.New("task name is required")
    }
    // Add more validation as needed
    return nil
}
