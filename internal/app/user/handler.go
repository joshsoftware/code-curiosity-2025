package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

func UpdateUserEmail(userService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var requestBody Email
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
			response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
			return
		}

		err = userService.UpdateUserEmail(ctx, requestBody.Email)
		if err != nil {
			slog.Error("failed to update user email", "error", err)
			status, errorMessage := apperrors.MapError(err)
			response.WriteJson(w, status, errorMessage, nil)
			return
		}

		response.WriteJson(w, http.StatusOK, "email updated successfully", nil)
	}
}
