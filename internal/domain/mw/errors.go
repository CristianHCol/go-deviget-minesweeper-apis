package mw

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// custom errors just for specifics, for general errors handles the original ones
var (
	ErrUserNotActiveOrDel    = status.New(codes.InvalidArgument, "[user]: you can not perform this operation with a non active or deleted user").Err()
	ErrUserKeyDuplicated     = status.New(codes.AlreadyExists, "[user]: this key (email or mobile) is alredy taken try another").Err()
	ErrNoUserAffected        = status.New(codes.DataLoss, "[user]: no user found to be processed, check the information sent and try again").Err()
	ErrorValidation          = status.New(codes.InvalidArgument, "[user]: the user did not pass the validation").Err()
	ErrorBcrypto             = status.New(codes.Internal, "[user]: error handling the users sensitive information").Err()
	ErrUserNotFound          = status.New(codes.NotFound, "[user]: user not found").Err()
	ErrUserInvalidCredential = status.New(codes.Unauthenticated, "[user]: invalid credentials").Err()
	ErrUserInvalidToken      = status.New(codes.Unauthenticated, "[user]: invalid authorization token provided").Err()
	ErrInternalUser          = status.New(codes.Internal, "[user]: an unexpected error has occurred trying to process the request").Err()
)

// Error codes the usage of this codes is to reply to to the clients
var (
	UserNoActive           = "USR_001"
	UserDuplicated         = "USR_002"
	UserNoAffected         = "USR_003"
	UserErrorValidation    = "USR_004"
	UserSensitiveInfo      = "USR_005"
	UserInternalError      = "USR_500"
	UserNotFound           = "USR_404"
	UserInvalidCredentials = "USR_006"
	InvalidToken           = "USR_007"
)

// HTTPErrorsCode map the service errors with specific http status code
var HTTPErrorsCode = map[string]int{
	"USR_001": http.StatusBadRequest,
	"USR_002": http.StatusConflict,
	"USR_003": http.StatusNotModified,
	"USR_004": http.StatusBadRequest,
	"USR_005": http.StatusInternalServerError,
	"USR_500": http.StatusInternalServerError,
	"USR_404": http.StatusNotFound,
	"USR_006": http.StatusUnauthorized,
	"USR_007": http.StatusUnauthorized,
}

// ServiceErrorsCode map the errors with specific service errors code
var ServiceErrorsCode = map[error]string{
	ErrUserNotActiveOrDel:    "USR_001",
	ErrUserKeyDuplicated:     "USR_002",
	ErrNoUserAffected:        "USR_003",
	ErrorValidation:          "USR_004",
	ErrorBcrypto:             "USR_005",
	ErrUserNotFound:          "USR_404",
	ErrUserInvalidCredential: "USR_006",
	ErrUserInvalidToken:      "USR_007",
	ErrInternalUser:          "USR_500",
}
