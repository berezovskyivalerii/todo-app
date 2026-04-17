package tasksservice

import (
	"context"

	"github.com/berezovskyivalerii/todo-app/internal/core/domain"
)

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, userID *int, limit *int, offset *int) ([]domain.Task, error)
	GetTask(ctx context.Context, id int) (domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	PatchTask(ctx context.Context, id int, patch domain.Task) (domain.Task, error)
}

type TasksService struct {
	tasksRepository TasksRepository
}

func NewTasksService(
	tasksRepository TasksRepository,
) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
