package mw

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// custom errors just for specifics, for general errors handles the original ones
var (
	ErrUserKeyDuplicated = status.New(codes.AlreadyExists, "[user]: this username is alredy taken try another").Err()
	ErrUserNotFound      = status.New(codes.NotFound, "[user]: username not found, please create one!").Err()
	ErrInternalUser      = status.New(codes.Internal, "[user]: an unexpected error has occurred trying to process the request").Err()
)

// Error codes the usage of this codes is to reply to to the clients
var (
	UserDuplicated = "USR_001"
	UserNotFound   = "USR_404"
	InternalError  = "USR_500"
)

// HTTPErrorsCode map the service errors with specific http status code
var HTTPErrorsCode = map[string]int{
	"USR_001": http.StatusBadRequest,
	"USR_404": http.StatusNotFound,
	"USR_500": http.StatusInternalServerError,
}

// ServiceErrorsCode map the errors with specific service errors code
var ServiceErrorsCode = map[error]string{
	ErrUserKeyDuplicated: "USR_001",
	ErrUserNotFound:      "USR_404",
	ErrInternalUser:      "USR_500",
}
