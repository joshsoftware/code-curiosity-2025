package utils

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

func FormatIntSliceForQuery(ids []int) string {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = fmt.Sprintf("%d", id)
	}

	return strings.Join(strIDs, ",")
}

func DoGet(httpClient *http.Client, url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		slog.Error("failed to create GET request", "error", err)
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		slog.Error("failed to send GET request", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error reading body", "error", err)
		return nil, err
	}

	return body, nil
}

func ValidateYearQueryParam(yearVal string) (int, error) {
	year, err := strconv.Atoi(yearVal)
	if err != nil {
		slog.Error("error converting year string value to int")
		return 0, err
	}

	if year < 2025 || year > time.Now().Year() {
		slog.Error("invalid year value")
		return 0, apperrors.ErrInvalidQueryParams
	}

	return year, nil
}

func ValidateMonthQueryParam(monthVal string) (int, error) {
	month, err := strconv.Atoi(monthVal)
	if err != nil {
		slog.Error("error converting month string value to int")
		return 0, err
	}

	if month < 0 || month > 12 {
		slog.Error("invalid month value")
		return 0, apperrors.ErrInvalidQueryParams
	}

	return month, nil
}
