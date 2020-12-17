package common

import (
	"net/http"
)

type (
	Error interface {
		error
		StatusCode() int
	}

	dataAlreadyExist  string
	dataNotFound      string
	paramsAreNeeded   string
	paramsAreWrong    string
	systemError       string
	operateNotAllowed string
	operateFailure    string
)

var (
	DataAlreadyExist  dataAlreadyExist  = "数据已存在"
	DataNotFound      dataNotFound      = "数据不存在"
	ParamsAreNeeded   paramsAreNeeded   = "缺乏参数"
	ParamsAreWrong    paramsAreWrong    = "参数错误"
	SystemError       systemError       = "系统错误"
	OperateNotAllowed operateNotAllowed = "禁止操作"
	OperateFailure    operateFailure    = "操作失败"
)

func (err *dataAlreadyExist) Error() string {
	return string(*err)
}
func (err *dataNotFound) Error() string {
	return string(*err)
}
func (err *paramsAreNeeded) Error() string {
	return string(*err)
}
func (err *paramsAreWrong) Error() string {
	return string(*err)
}
func (err *systemError) Error() string {
	return string(*err)
}
func (err *operateNotAllowed) Error() string {
	return string(*err)
}
func (err *operateFailure) Error() string {
	return string(*err)
}

func (err *dataAlreadyExist) New(info string) *dataAlreadyExist {
	if info == "" {
		return err
	}
	e := dataAlreadyExist(info)
	return &e
}

func (err *dataNotFound) New(info string) *dataNotFound {
	if info == "" {
		return err
	}
	e := dataNotFound(info)
	return &e
}

func (err *paramsAreNeeded) New(info string) *paramsAreNeeded {
	if info == "" {
		return err
	}
	e := paramsAreNeeded(info)
	return &e
}

func (err *paramsAreWrong) New(info string) *paramsAreWrong {
	if info == "" {
		return err
	}
	e := paramsAreWrong(info)
	return &e
}

func (err *systemError) New(info string) *systemError {
	if info == "" {
		return err
	}
	e := systemError(info)
	return &e
}

func (err *operateNotAllowed) New(info string) *operateNotAllowed {
	if info == "" {
		return err
	}
	e := operateNotAllowed(info)
	return &e
}

func (err *operateFailure) New(info string) *operateFailure {
	if info == "" {
		return err
	}
	e := operateFailure(info)
	return &e
}

func (err *dataAlreadyExist) StatusCode() int {
	return http.StatusForbidden
}
func (err *dataNotFound) StatusCode() int {
	return http.StatusNotFound
}
func (err *paramsAreNeeded) StatusCode() int {
	return http.StatusBadRequest
}
func (err *paramsAreWrong) StatusCode() int {
	return http.StatusBadRequest
}
func (err *systemError) StatusCode() int {
	return http.StatusInternalServerError
}
func (err *operateNotAllowed) StatusCode() int {
	return http.StatusForbidden
}
func (err *operateFailure) StatusCode() int {
	return http.StatusNotModified
}
