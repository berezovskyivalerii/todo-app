package taskshttp

import (
	"net/http"

	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	corerequest "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/request"
	coreresponse "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/response"
)

type DeleteTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := corelogger.FromContext(ctx)
	responseHandler := coreresponse.NewHTTPResponseHandler(logger, rw)

	taskID, err := corerequest.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get taskID path value")
		return
	}

	err = h.tasksService.DeleteTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to delete task")
		return
	}

	responseHandler.NoContentResponse()
}
