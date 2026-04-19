package statisticsrepository

import corepostgrespool "github.com/berezovskyivalerii/todo-app/internal/core/repository/postgres/pool"

type StatisticsRepository struct {
	pool corepostgrespool.Pool
}

func NewStatiscticsRepository(
	pool corepostgrespool.Pool,
) *StatisticsRepository {
	return &StatisticsRepository{
		pool: pool,
	}
}
