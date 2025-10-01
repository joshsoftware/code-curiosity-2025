package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/jwt"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type txKeyType struct{}

var txKey = txKeyType{}

type contextKey string

const (
	UserIdKey    contextKey = "userId"
	IsBlockedKey contextKey = "isBlocked"
	IsAdminKey   contextKey = "isAdmin"
)

func EmbedTxInContext(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func ExtractTxFromContext(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(txKey).(*sqlx.Tx)
	return tx, ok
}

func CorsMiddleware(next http.Handler, appCfg config.AppConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", appCfg.ClientURL)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Authentication(next http.HandlerFunc, appCfg config.AppConfig) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.WriteJson(w, http.StatusUnauthorized, apperrors.ErrAuthorizationFailed.Error(), nil)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseJWT(tokenString, appCfg)
		if err != nil {
			response.WriteJson(w, http.StatusUnauthorized, apperrors.ErrAuthorizationFailed.Error(), nil)
			return
		}

		userId := token.UserId
		ctx := context.WithValue(r.Context(), UserIdKey, userId)
		isAdmin := token.IsAdmin
		ctx = context.WithValue(ctx, IsAdminKey, isAdmin)
		IsBlocked := token.IsBlocked
		ctx = context.WithValue(ctx, IsBlockedKey, IsBlocked)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func AuthorizeUnblockedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		isBlocked, ok := ctx.Value(IsBlockedKey).(bool)
		if !ok {
			response.WriteJson(w, http.StatusInternalServerError, apperrors.ErrContextValue.Error(), nil)
			return
		}

		if isBlocked {
			response.WriteJson(w, http.StatusForbidden, apperrors.ErrUserBlocked.Error(), nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthorizeAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		isAdmin, ok := ctx.Value(IsAdminKey).(bool)
		if !ok {
			response.WriteJson(w, http.StatusInternalServerError, apperrors.ErrContextValue.Error(), nil)
			return
		}

		if !isAdmin {
			response.WriteJson(w, http.StatusUnauthorized, apperrors.ErrUnauthorizedAccess.Error(), nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
