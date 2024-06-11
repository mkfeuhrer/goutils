package errors

import (
	"encoding/json"
	"errors"
	"fmt"

	pkgErrors "github.com/pkg/errors"
)

// Error represents a structured error with a unique code, message, additional data, and an underlying error.
// It is used to provide a consistent error handling mechanism across the application.
type Error struct {
	// Code is a unique string that identifies the type of error.
	// This code can be used programmatically to handle specific error cases.
	Code string `json:"code"`

	// StatusCode denotes http status code. Useful in API Error handling
	StatusCode int `json:"status_code,omitempty"`

	// Message is a human-readable description of the error.
	// This message is intended to be sent to the client to provide details about the error.
	Message string `json:"message"`

	// Data is a map containing additional contextual information about the error.
	// This can include any relevant key-value pairs that provide more details about the error situation.
	Data map[string]interface{} `json:"data,omitempty"`

	// Err is the underlying error that triggered this Error struct.
	// It can be used to trace back to the original error for debugging purposes.
	Err error `json:"-"`
}

func (err *Error) Error() string {
	if err.Err != nil {
		return fmt.Sprintf("code: %s, message: %s, underlying error: %v", err.Code, err.Message, err.Err)
	}
	return fmt.Sprintf("code: %s, message: %s", err.Code, err.Message)
}

// GetCode returns the error code.
func (err *Error) GetCode() string {
	return err.Code
}

// GetData returns the data map of the error.
func (err *Error) GetData() map[string]interface{} {
	return err.Data
}

// GetError returns the underlying error.
func (err *Error) GetError() error {
	return err.Err
}

// Bytes returns the json byte array. Useful for sending in HTTP response
func (err *Error) Bytes() []byte {
	out, _ := json.Marshal(err)
	return out
}

// NewError creates a new Error instance with the provided code, message, data, and an optional underlying error.
func NewError(code, message string, data map[string]interface{}, underlyingErr error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
		Err:     underlyingErr,
	}
}

// NewAPIError create a new Error instance with provided code and http status code.
func NewAPIError(code string, statusCode int) *Error {
	return &Error{
		Code:       code,
		StatusCode: statusCode,
	}
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Wrap returns new error by annotating the passed error with message
func Wrap(err error, message string, obj ...interface{}) error {
	if obj == nil || len(obj) == 0 {
		return pkgErrors.Wrap(err, message)
	} else {
		return pkgErrors.Wrap(err, fmt.Sprintf(message, obj...))
	}
}

// Cause returns the underlying cause of the error by unwrapping the error
// StackTrace.
func Cause(err error) error {
	return pkgErrors.Cause(err)
}
