package common

import (
	"errors"
	"net/http"
)

// custom errors just for specifics, fro general errors handles the original ones
var (
	ErrUserUnauthorized   = errors.New("[game]: user does not exists")
	ErrOperationForbidden = errors.New("[game]: game does not exist")
)

// Error codes the usage of this codes is to reply to to the clients
var (
	UserNoAuthorized   = "GM_001"
	OperationForbidden = "GM_002"
)

// HTTPErrorsCode map the service errors with specific http status code
var HTTPErrorsCode = map[string]int{
	"ERR_500": http.StatusInternalServerError,
	"ERR_001": http.StatusUnauthorized,
	"ERR_002": http.StatusForbidden,
}

// ServiceErrorsCode map the errors with specific service errors code
var ServiceErrorsCode = map[error]string{
	ErrUserUnauthorized:   "ERR_001",
	ErrOperationForbidden: "ERR_002",
}
