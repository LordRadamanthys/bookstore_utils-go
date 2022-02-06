package rest_errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func BadRequestError(message string, err error) *RestErr {

	result := &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}

func NotFoundError(message string, err error) *RestErr {

	result := &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}

func InternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}

func UnauthorizedError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}
