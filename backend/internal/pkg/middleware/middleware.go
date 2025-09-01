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
	UserIdKey  contextKey = "userId"
	IsAdminKey contextKey = "isAdmin"
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
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
