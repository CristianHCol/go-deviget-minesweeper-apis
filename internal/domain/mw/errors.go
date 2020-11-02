package mw

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// custom errors just for specifics, for general errors handles the original ones
var (
	ErrUserKeyDuplicated    = status.New(codes.AlreadyExists, "[user]: this username is alredy taken try another").Err()
	ErrUserNotFound         = status.New(codes.NotFound, "[user]: username not found, please create one!").Err()
	ErrInternalError        = status.New(codes.Internal, "[user]: an unexpected error has occurred trying to process the request").Err()
	ErrGameNotFound         = status.New(codes.NotFound, "[game]: game not found, please create one!").Err()
	ErrGamePermissionDenied = status.New(codes.PermissionDenied, "[game]: the username does not have permissions for this game").Err()
	ErrGameAlreadyStarted   = status.New(codes.PermissionDenied, "[game]: the game is already started").Err()
	ErrGameIsNotStarted     = status.New(codes.PermissionDenied, "[game]: the game is not started or is over").Err()
	ErrGameIsOver           = status.New(codes.OK, "[game]: the game is over!!").Err()
	ErrGameIsWon            = status.New(codes.OK, "[game]: You won!!").Err()
)

// Error codes the usage of this codes is to reply to to the clients
var (
	UserDuplicated           = "USR_001"
	UserNotFound             = "USR_404"
	InternalError            = "USR_500"
	GameNotFound             = "GME_404"
	GameGamePermissionDenied = "GME_403"
	GameAlreadyStarted       = "GME_001"
	GameIsNotStarted         = "GME_002"
	GameIsOver               = "GME_003"
	GameIsWon                = "GME_004"
)

// HTTPErrorsCode map the service errors with specific http status code
var HTTPErrorsCode = map[string]int{
	"USR_001": http.StatusBadRequest,
	"USR_404": http.StatusNotFound,
	"USR_500": http.StatusInternalServerError,
	"GME_404": http.StatusNotFound,
	"GME_403": http.StatusForbidden,
	"GME_001": http.StatusBadRequest,
	"GME_002": http.StatusBadRequest,
	"GME_003": http.StatusOK,
	"GME_004": http.StatusOK,
}

// ServiceErrorsCode map the errors with specific service errors code
var ServiceErrorsCode = map[error]string{
	ErrUserKeyDuplicated:    "USR_001",
	ErrUserNotFound:         "USR_404",
	ErrInternalError:        "USR_500",
	ErrGameNotFound:         "GME_404",
	ErrGamePermissionDenied: "GME_403",
	ErrGameAlreadyStarted:   "GME_001",
	ErrGameIsNotStarted:     "GME_002",
	ErrGameIsOver:           "GME_003",
	ErrGameIsWon:            "GME_004",
}
