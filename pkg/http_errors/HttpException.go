package http_errors

import (
	"go-clean-arch/pkg/utils"
	"net/http"
	"time"
)

// HttpException struct
type HttpException struct {
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHttpException(status int, message string) *HttpException {
	return &HttpException{
		Message: message,
		Status:  status,
	}
}

func InternalServerError(err error) *HttpException {
	return NewHttpException(http.StatusInternalServerError, utils.NewValidatorError(err))
}

func BadRequestError(err error) *HttpException {
	return NewHttpException(http.StatusBadRequest, utils.NewValidatorError(err))
}

func NotFoundError(err error) *HttpException {
	return NewHttpException(http.StatusNotFound, utils.NewValidatorError(err))
}
