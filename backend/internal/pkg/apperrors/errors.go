package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServer = errors.New("internal server error")

	ErrInvalidRequestBody = errors.New("invalid or missing parameters in the request body")
	ErrInvalidQueryParams = errors.New("invalid or missing query parameters")
	ErrFailedMarshal      = errors.New("failed to parse request body")

	ErrUnauthorizedAccess = errors.New("unauthorized. please provide a valid access token")
	ErrAccessForbidden    = errors.New("access forbidden")
	ErrInvalidToken       = errors.New("invalid or expired token")

	ErrFailedInitializingLogger = errors.New("failed to initialize logger")
	ErrNoAppConfigPath          = errors.New("no config path provided")
	ErrFailedToLoadAppConfig    = errors.New("failed to load environment configuration")

	ErrLoginWithGithubFailed    = errors.New("failed to login with Github")
	ErrGithubTokenExchangeFailed = errors.New("failed to exchange Github token")
	ErrFailedToGetGithubUser = errors.New("failed to get Github user info")
	ErrFailedToGetUserEmail = errors.New("failed to get user email from Github")

	ErrUserNotFound = errors.New("user not found")
	ErrUserCreationFailed = errors.New("failed to create user")

	ErrJWTCreationFailed = errors.New("failed to create jwt token")
	ErrAuthorizationFailed=errors.New("failed to authorize user")
)

func MapError(err error) (statusCode int, errMessage string) {
	switch err {
	case ErrInvalidRequestBody, ErrInvalidQueryParams:
		return http.StatusBadRequest, err.Error()
	case ErrUnauthorizedAccess:
		return http.StatusUnauthorized, err.Error()
	case ErrAccessForbidden:
		return http.StatusForbidden, err.Error()
	case ErrUserNotFound:
		return http.StatusNotFound, err.Error()
	case ErrInvalidToken:
		return http.StatusUnprocessableEntity, err.Error()
	default:
		return http.StatusInternalServerError, ErrInternalServer.Error()
	}
}
