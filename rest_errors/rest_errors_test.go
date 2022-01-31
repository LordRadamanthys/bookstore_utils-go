package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {

	err := InternalServerError("this is the message", errors.New("teste error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "teste error", err.Causes[0])

}

func TestBadRequestError(t *testing.T) {
	err := BadRequestError("this is the message",
		errors.New("teste error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "teste error", err.Causes[0])
}

func TestNotFoundError(t *testing.T) {
	err := NotFoundError("this is the message", errors.New("teste error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "not_found", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "teste error", err.Causes[0])
}
