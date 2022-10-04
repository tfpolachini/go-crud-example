package controller

import (
	"errors"
	"net/http"

	"github.com/tfpolachini/go-crud-example/domain/service"
)

func HttpStatusOf(err error) int {
	if errors.Is(err, service.ErrResourceNotFound) {
		return http.StatusNotFound
	} else if errors.Is(err, service.ErrResourceAlreadyExists) {
		return http.StatusConflict
	} else if errors.Is(err, service.ErrInvalidAction) {
		return http.StatusBadRequest
	} else {
		return http.StatusInternalServerError
	}
}
