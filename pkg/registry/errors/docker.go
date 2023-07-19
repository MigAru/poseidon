package errors

import "fmt"

type ErrorResponse struct {
	Errors []error `json:"errors"`
}

func CreateErrorResponse(err error) ErrorResponse {
	return ErrorResponse{[]error{err}}
}

type DockerError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e DockerError) Error() string {
	return fmt.Sprintf("")
}
