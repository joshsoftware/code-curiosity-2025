package utils

import (
	"fmt"
	"strings"
)

func FormatIntSliceForQuery(ids []int) string {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = fmt.Sprintf("%d", id)
	}

	return strings.Join(strIDs, ",")
}
