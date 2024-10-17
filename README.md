# practice_layers

---

## Project Overview

This project consists of multiple layers that are cleanly decoupled through the use of **interfaces**. The layers are:

1. **API Handler Layer**: Handles HTTP requests and responses.
2. **Service Layer**: Implements the core business logic, including validation.
3. **Repository Layer**: Manages data storage and retrieval, interacting with PostgreSQL.
4. **Validation Layer**: Contains logic to validate data before it reaches the repository.

Each layer communicates with the others through well-defined interfaces, ensuring that the implementation details are abstracted away, making the system easier to extend, test, and maintain.

### 1. API Handler Layer

**File: `handlers/task_handler.go`**

This layer is responsible for handling HTTP requests (GET and POST). When a request is received, it passes the data down to the service layer for processing.

- **Decoupling**: The handler doesn’t know the implementation details of the service. It only relies on the service interface (`TaskService`), allowing you to easily swap the underlying logic without affecting the handler.
  
- **Dependency Injection (DI)**: The `TaskService` and repository are injected into the handler function as dependencies. This ensures that the handler isn't tightly coupled to specific implementations.

```go
func TaskHandler(w http.ResponseWriter, r *http.Request) {
    repo := repository.NewPostgresTaskRepository(db) // Repository created here
    validator := validation.NewTaskValidator()       // Validation service created here
    service := services.NewTaskService(repo, validator) // Service injected with repository and validator
    ...
}
```

### 2. Service Layer

**File: `services/task_service.go`**

The service layer contains the core logic for your application. It is responsible for handling business rules like task creation and retrieval. It also manages the validation logic by invoking the validation layer.

- **Decoupling**: The service layer uses interfaces for both the repository and validation layers. This allows the service to focus on business logic without needing to know the specifics of how data is fetched or validated.

- **Dependency Injection**: The repository and validation services are injected into the service when it's initialized. This keeps the service flexible, as you can provide different implementations of these dependencies (e.g., switch from a mock repository to a PostgreSQL repository for testing vs. production).

```go
type TaskService struct {
    repo      interfaces.TaskRepository       // Injected dependency (Repository Layer)
    validator interfaces.ValidationService    // Injected dependency (Validation Layer)
}

func NewTaskService(repo interfaces.TaskRepository, validator interfaces.ValidationService) interfaces.TaskService {
    return &TaskService{repo: repo, validator: validator}
}

func (s *TaskService) CreateTask(task schemas.Task) (schemas.Task, error) {
    err := s.validator.ValidateTask(task) // Validation before creating the task
    if err != nil {
        return schemas.Task{}, err
    }
    return s.repo.SaveTask(task), nil
}
```

### 3. Repository Layer

**File: `repository/task_repository.go`**

This layer manages data storage and retrieval from the PostgreSQL database. It implements the `TaskRepository` interface, which is used by the service layer.

- **Decoupling**: The service layer only interacts with the repository through the `TaskRepository` interface. This makes it easy to replace or mock the repository for testing or future enhancements (like changing databases).

```go
type PostgresTaskRepository struct {
    db *sql.DB // Database connection injected
}

func NewPostgresTaskRepository(db *sql.DB) interfaces.TaskRepository {
    return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) SaveTask(task schemas.Task) schemas.Task {
    // Logic for saving task to the PostgreSQL database
}
```

### 4. Validation Layer

**File: `validation/task_validator.go`**

The validation layer ensures that the data adheres to business rules before reaching the repository. In this case, it checks if the task's required fields are present.

- **Decoupling**: The service layer relies on the `ValidationService` interface. The actual validation logic is hidden behind this interface, meaning the service doesn’t need to know the specifics of how validation is performed.
  
- **Dependency Injection**: The validation layer is injected into the service layer, allowing you to easily change or mock validation logic as needed.

```go
type TaskValidator struct{}

func NewTaskValidator() interfaces.ValidationService {
    return &TaskValidator{}
}

func (v *TaskValidator) ValidateTask(task schemas.Task) error {
    if task.Name == "" {
        return errors.New("Task name cannot be empty")
    }
    // Other validation logic
    return nil
}
```

### Interfaces and Decoupling

The key to decoupling in this project lies in the use of **interfaces**. Interfaces allow you to define a contract for what a layer should do, without dictating how it does it.

- **Repository Interface**: Defines the methods that any repository (e.g., PostgreSQL, in-memory) must implement.
- **Service Interface**: Ensures the service has a standard way of interacting with the handler layer.
- **Validation Interface**: Abstracts the validation logic, making it interchangeable.

```go
package interfaces

import "taskmanager/schemas"

type TaskRepository interface {
    FetchTasks() []schemas.Task
    SaveTask(task schemas.Task) schemas.Task
}

type TaskService interface {
    GetTasks() []schemas.Task
    CreateTask(task schemas.Task) (schemas.Task, error)
}

type ValidationService interface {
    ValidateTask(task schemas.Task) error
}
```

### Dependency Injection (DI)

Dependency injection is used throughout the project to provide flexibility and testability. Instead of hardcoding dependencies (e.g., creating a new repository directly inside the service), they are passed in as arguments. This makes it easier to swap out dependencies (e.g., use a mock repository for testing).

- **Handler Layer**: Injects the `TaskService`, which in turn has the `TaskRepository` and `ValidationService`.
- **Service Layer**: Injects the `TaskRepository` (to interact with the database) and `ValidationService` (to validate the task).

### Database Integration (PostgreSQL)

The project connects to PostgreSQL using `database/sql` and the `pgx` driver. The `repository` layer handles all database operations. This setup ensures that the service layer doesn’t need to know how the data is stored or fetched.

**Database Connection Setup** (`config.ConnectDB()`):
- Connection to the PostgreSQL database is established using environment variables (loaded via `godotenv`).
- The repository is initialized with this database connection, ensuring that all interactions with the database are routed through the repository.

```go
db := config.ConnectDB()
repo := repository.NewPostgresTaskRepository(db) // Repository initialized with db connection
```

### Conclusion

This project uses interfaces and dependency injection to achieve clean separation of concerns between layers. Each layer focuses on a specific responsibility (handling requests, business logic, validation, and database operations) and interacts with the others through interfaces, keeping the system modular and flexible.

This approach makes your code easier to test, extend, and maintain while adhering to SOLID principles, especially the **Dependency Inversion Principle** (DIP), which is achieved by injecting dependencies through interfaces rather than concrete implementations.

---