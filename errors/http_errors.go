package errors

import "net/http"

var (
	InternalServerErr = NewAPIError("INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
	BadRequestErr     = NewAPIError("BAD_REQUEST", http.StatusBadRequest)
	UnAuthorizedErr   = NewAPIError("UNAUTHORIZED", http.StatusUnauthorized)
)
