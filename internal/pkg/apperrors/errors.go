package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrContextValue   = errors.New("error obtaining value from context")
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

	ErrLoginWithGithubFailed     = errors.New("failed to login with Github")
	ErrGithubTokenExchangeFailed = errors.New("failed to exchange Github token")
	ErrFailedToGetGithubUser     = errors.New("failed to get Github user info")
	ErrFailedToGetUserEmail      = errors.New("failed to get user email from Github")

	ErrUserNotFound       = errors.New("user not found")
	ErrUserCreationFailed = errors.New("failed to create user")

	ErrJWTCreationFailed   = errors.New("failed to create jwt token")
	ErrAuthorizationFailed = errors.New("failed to authorize user")

	ErrRepoNotFound                    = errors.New("repository not found")
	ErrRepoCreationFailed              = errors.New("failed to create repo for user")
	ErrCalculatingUserRepoTotalCoins   = errors.New("error calculating total coins earned by user for the repository")
	ErrFetchingUsersContributedRepos   = errors.New("error fetching users contributed repositories")
	ErrFetchingUserContributionsInRepo = errors.New("error fetching users contribution in repository")

	ErrFetchingFromBigquery        = errors.New("error fetching contributions from bigquery service")
	ErrNextContribution            = errors.New("error while loading next bigquery contribution")
	ErrContributionCreationFailed  = errors.New("failed to create contrbitution")
	ErrFetchingRecentContributions = errors.New("failed to fetch users five recent contributions")
	ErrFetchingAllContributions    = errors.New("failed to fetch all contributions for user")
	ErrContributionScoreNotFound   = errors.New("failed to get contributionscore details for given contribution type")
	ErrFetchingContribution        = errors.New("error fetching contribution by github repo id")
	ErrContributionNotFound        = errors.New("contribution not found")

	ErrTransactionCreationFailed = errors.New("error failed to create transaction")
	ErrTransactionNotFound       = errors.New("error transaction for the contribution id does not exist")
)

func MapError(err error) (statusCode int, errMessage string) {
	switch err {
	case ErrInvalidRequestBody, ErrInvalidQueryParams, ErrContextValue:
		return http.StatusBadRequest, err.Error()
	case ErrUnauthorizedAccess:
		return http.StatusUnauthorized, err.Error()
	case ErrAccessForbidden:
		return http.StatusForbidden, err.Error()
	case ErrUserNotFound, ErrRepoNotFound, ErrContributionNotFound:
		return http.StatusNotFound, err.Error()
	case ErrInvalidToken:
		return http.StatusUnprocessableEntity, err.Error()
	default:
		return http.StatusInternalServerError, ErrInternalServer.Error()
	}
}
