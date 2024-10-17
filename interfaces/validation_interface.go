package interfaces

import "taskmanager/schemas"

type ValidationService interface {
    ValidateTask(task schemas.Task) error
}
