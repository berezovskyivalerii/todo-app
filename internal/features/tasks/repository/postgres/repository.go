package tasksrepository

import corepostgrespool "github.com/berezovskyivalerii/todo-app/internal/core/repository/postgres/pool"

type TasksRepository struct {
	pool corepostgrespool.Pool
}

func NewTasksRepository(
	pool corepostgrespool.Pool,
) *TasksRepository {
	return &TasksRepository{
		pool: pool,
	}
}
