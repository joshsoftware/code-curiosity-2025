package utils

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
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
