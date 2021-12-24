package errors

import "net/http"

var (
	ERR_UNKNOWN        = New(http.StatusInternalServerError, 0, "unknown error")
	ERR_USER_NOT_FOUND = New(http.StatusBadRequest, 1, "user not found")
	ERR_TEAM_NOT_FOUND = New(http.StatusBadRequest, 2, "team not found")
)
