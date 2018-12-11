package gerrors

import (
	"encoding/json"
	"net/http"
)

type GlobalError struct {
	Code        int
	ServiceName string
	Message     string
	Detail      interface{}
	StatusCode  int
}

func (ge *GlobalError) Error() string {
	b, _ := json.Marshal(ge)
	return string(b)
}

func New(code int, statusCode int, message string, detail interface{}) error {
	if message=="" {
		message = http.StatusText(statusCode)
	}
	return &GlobalError{
		Code: code,
		// ServiceName: serviceName,
		Message:    message,
		Detail:     detail,
		StatusCode: statusCode,
	}
}

// BadRequest generates a 400 error.
func BadRequest(code int, message string, detail interface{}) error {
	return &GlobalError{
		Code: code,
		// ServiceName: serviceName,
		StatusCode: 400,
		Message:    message,
		Detail:     detail,
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(code int, message string, detail interface{}) error {
	return &GlobalError{
		Code: code,
		// ServiceName: serviceName,
		Message:    message,
		StatusCode: 401,
		Detail:     detail,
	}
}

// Forbidden generates a 403 error.
func Forbidden(code int, message string, detail ...interface{}) error {
	if message == "" {
		message = http.StatusText(http.StatusForbidden)
	}
	return &GlobalError{
		Code: code,
		// ServiceName: serviceName,
		Message:    message,
		StatusCode: http.StatusForbidden,
		Detail:     detail,
	}
}

// NotFound generates a 404 error.
func NotFound(code int, message string, detail ...interface{}) error {
	if message == "" {
		message = http.StatusText(http.StatusNotFound)
	}
	return &GlobalError{
		Code: code,
		// ServiceName: serviceName,
		Message:    message,
		StatusCode: http.StatusNotFound,
		Detail:     detail,
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(code int, message string, detail ...interface{}) error {
	if message == "" {
		message = http.StatusText(http.StatusInternalServerError)
	}
	return &GlobalError{
		Code: code,
		// ServiceName: serviceName,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Detail:     detail,
	}
}

// Conflict generates a 409 error.
func Conflict(code int, message string, detail ...interface{}) error {
	if message == "" {
		message = http.StatusText(http.StatusConflict)
	}
	return &GlobalError{
		Code: code,
		// ServiceName: serviceName,
		Message:    message,
		StatusCode: http.StatusConflict,
		Detail:     detail,
	}
}

type ValidateError map[string][]string

// UnprocessableEntity generates a 422 error.
func UnprocessableEntity(code int, ve ValidateError) error {
	return &GlobalError{
		Code: code,
		// ServiceName: serviceName,
		Message:    "The given data failed to pass validation.",
		StatusCode: http.StatusUnprocessableEntity,
		Detail:     ve,
	}
}
