// Package coreresponse contains reponse helpers
package coreresponse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	coreerrors "github.com/berezovskyivalerii/todo-app/internal/core/errors"
	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *corelogger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(log *corelogger.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{log: log, rw: rw}
}

func (h *HTTPResponseHandler) JSONResponse(responseBody any, statusCode int) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		h.log.Error("write HTTP reponse", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.rw.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, coreerrors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, coreerrors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	case errors.Is(err, coreerrors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) PanicReponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) errorResponse(statusCode int, err error, msg string) {
	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	h.JSONResponse(response, statusCode)
}
