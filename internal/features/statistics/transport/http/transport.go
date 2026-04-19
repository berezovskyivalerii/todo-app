package statisticshttp

import (
	"context"
	"net/http"
	"time"

	"github.com/berezovskyivalerii/todo-app/internal/core/domain"
	coreserver "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(ctx context.Context, userID *int, from *time.Time, to *time.Time) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(
	statisticsService StatisticsService,
) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []coreserver.Route {
	return []coreserver.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
