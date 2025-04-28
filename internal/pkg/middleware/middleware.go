package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	constant "github.com/joshsoftware/code-curiosity-2025/internal/pkg/constants"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/jwt"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appCfg := config.GetAppConfig()
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

// Checks whether the user is valid, that is signed in before the user can access the next handler functionality
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.WriteJson(w, http.StatusUnauthorized, apperrors.ErrAuthorizationFailed.Error(), nil)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseJWT(tokenString)
		if err != nil {
			response.WriteJson(w, http.StatusUnauthorized, apperrors.ErrAuthorizationFailed.Error(), nil)
			return
		}

		userId := token.UserId
		ctx := context.WithValue(r.Context(),constant.UserIdKey, userId)
		isAdmin := token.IsAdmin
		ctx = context.WithValue(ctx, constant.IsAdminKey, isAdmin)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
