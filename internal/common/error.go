package common

import "net/http"

type Error interface {
	error
	StatusCode() int
}

type (
	DataAlreadyExist  string
	DataNotFound      string
	ParamsAreNeeded   string
	ParamsAreWrong    string
	SystemError       string
	OperateNotAllowed string
	OperateFailure    string
)

func (err *DataAlreadyExist) Error() string {
	return string(*err)
}
func (err *DataNotFound) Error() string {
	return string(*err)
}
func (err *ParamsAreNeeded) Error() string {
	return string(*err)
}
func (err *ParamsAreWrong) Error() string {
	return string(*err)
}
func (err *SystemError) Error() string {
	return string(*err)
}
func (err *OperateNotAllowed) Error() string {
	return string(*err)
}
func (err *OperateFailure) Error() string {
	return string(*err)
}

func (err *DataAlreadyExist) StatusCode() int {
	return http.StatusForbidden
}
func (err *DataNotFound) StatusCode() int {
	return http.StatusNotFound
}
func (err *ParamsAreNeeded) StatusCode() int {
	return http.StatusBadRequest
}
func (err *ParamsAreWrong) StatusCode() int {
	return http.StatusBadRequest
}
func (err *SystemError) StatusCode() int {
	return http.StatusInternalServerError
}
func (err *OperateNotAllowed) StatusCode() int {
	return http.StatusForbidden
}
func (err *OperateFailure) StatusCode() int {
	return http.StatusNotModified
}
