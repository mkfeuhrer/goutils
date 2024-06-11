package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	data := map[string]interface{}{
		"key": "value",
	}
	underlyingErr := errors.New("underlying error")

	err := NewError("CODE123", "An error occurred", data, underlyingErr)

	assert.Equal(t, "CODE123", err.Code)
	assert.Equal(t, "An error occurred", err.Message)
	assert.Equal(t, data, err.Data)
	assert.Equal(t, underlyingErr, err.Err)
}

func TestNewAPIError(t *testing.T) {
	err := NewAPIError("CODE404", http.StatusBadRequest)

	assert.Equal(t, "CODE404", err.Code)
	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "", err.Message)
	assert.Nil(t, err.Data)
	assert.Nil(t, err.Err)
}

func TestErrorMethod(t *testing.T) {
	data := map[string]interface{}{
		"key": "value",
	}
	underlyingErr := errors.New("underlying error")

	err := NewError("CODE123", "An error occurred", data, underlyingErr)

	expectedErrMsg := "code: CODE123, message: An error occurred, underlying error: underlying error"
	assert.Equal(t, expectedErrMsg, err.Error())

	errNoUnderlying := NewError("CODE124", "Another error occurred", nil, nil)
	expectedErrMsgNoUnderlying := "code: CODE124, message: Another error occurred"
	assert.Equal(t, expectedErrMsgNoUnderlying, errNoUnderlying.Error())
}

func TestGetCode(t *testing.T) {
	err := NewError("CODE123", "An error occurred", nil, nil)
	assert.Equal(t, "CODE123", err.GetCode())
}

func TestGetData(t *testing.T) {
	data := map[string]interface{}{
		"key": "value",
	}
	err := NewError("CODE123", "An error occurred", data, nil)
	assert.Equal(t, data, err.GetData())
}

func TestGetError(t *testing.T) {
	underlyingErr := errors.New("underlying error")
	err := NewError("CODE123", "An error occurred", nil, underlyingErr)
	assert.Equal(t, underlyingErr, err.GetError())
}

func TestBytes(t *testing.T) {
	data := map[string]interface{}{
		"key": "value",
	}
	err := NewError("CODE123", "An error occurred", data, nil)

	jsonBytes := err.Bytes()

	expectedJSON := `{"code":"CODE123","message":"An error occurred","data":{"key":"value"}}`
	assert.JSONEq(t, expectedJSON, string(jsonBytes))
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "wrapped message")

	assert.NotNil(t, wrappedErr)
	assert.Contains(t, wrappedErr.Error(), "original error")
	assert.Contains(t, wrappedErr.Error(), "wrapped message")
}

func TestWrapWithFormat(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "wrapped message: %s", "extra info")

	assert.NotNil(t, wrappedErr)
	assert.Contains(t, wrappedErr.Error(), "original error")
	assert.Contains(t, wrappedErr.Error(), "wrapped message: extra info")
}

func TestCause(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "wrapped message")

	assert.Equal(t, originalErr, Cause(wrappedErr))
}

func TestAs(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap(originalErr, "wrapped message")

	var targetErr *Error
	assert.False(t, As(originalErr, &targetErr))
	assert.False(t, As(wrappedErr, &targetErr))

	customErr := NewError("CODE123", "Custom error", nil, nil)
	assert.True(t, As(customErr, &targetErr))
	assert.Equal(t, customErr, targetErr)
}
