package errors

import (
	"fmt"
)

var (
	DefaultError = Define("default_error")
)

type Error struct {
	code       string
	message    string
	cause      error
	data       interface{}
	stacktrace frame
}

func (e Error) Error() string {
	return e.message
}

func Define(code string) Error {
	return Error{code: code}
}

func New(theType Error, cause error, message string) error {
	return newError(theType, cause, message, nil)
}

func NewWithData(theType Error, cause error, message string, data interface{}) error {
	return newError(theType, cause, message, data)
}

func Wrap(e error, message string) error {
	return wrap(e, message, nil)
}

func WrapWithData(e error, message string, data interface{}) error {
	return wrap(e, message, data)
}

func Is(e error, target error) bool {
	if err, ok := e.(Error); ok {
		if targetErr, ok := target.(Error); ok {
			return err.code == targetErr.code
		}
	}

	return e == target
}

func IsNot(e error, target error) bool {
	return !Is(e, target)
}

func Code(e error) string {
	if err, ok := e.(Error); ok {
		return err.code
	}

	return ""
}

func Data(e error) interface{} {
	if err, ok := e.(Error); ok {
		return err.data
	}

	return nil
}

func String(e error) string {
	if e == nil {
		return ""
	}

	if err, ok := e.(Error); ok {
		r := err.message + " | CODE: " + err.code

		if err.data != nil {
			r = fmt.Sprintf("%v | DATA: %+v", r, err.data)
		}

		r = fmt.Sprintf("%v | FILE: %v:%v", r, err.stacktrace.file(), err.stacktrace.line())

		if err.cause != nil {
			r = fmt.Sprintf("%v | CAUSE: {%v}", r, String(err.cause))
		}

		return r
	}

	return e.Error()
}

func newError(theType Error, cause error, message string, data interface{}) error {
	return Error{code: theType.code, message: message, cause: cause, data: data, stacktrace: stackLevel(2)}
}

func wrap(e error, message string, data interface{}) error {
	if err, ok := e.(Error); ok {
		return Error{code: err.code, message: message, cause: e, data: data, stacktrace: stackLevel(2)}
	}

	return Error{code: DefaultError.code, message: message, cause: e, data: data, stacktrace: stackLevel(2)}
}