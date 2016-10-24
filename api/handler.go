package api

import (
	"net/http"

	"github.com/anthonynsimon/parrot/errors"
	"github.com/anthonynsimon/parrot/render"
)

type apiHandlerFunc func(http.ResponseWriter, *http.Request) error

func (fn apiHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		respErr := errors.ErrInternal
		if castedErr, ok := err.(*errors.Error); ok {
			respErr = castedErr
		}
		render.JSONError(w, respErr)
	}
}

func ping(w http.ResponseWriter, r *http.Request) error {
	render.JSON(w, http.StatusOK, "Backend says hello.")
	return nil
}
