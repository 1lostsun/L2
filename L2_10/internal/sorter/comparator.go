package sorter

import (
	"strings"
)

func getKey(s string, col int, ignoreTrails bool) string {
	if ignoreTrails {
		strings.TrimRight(s, "\t")
	}

	parts := strings.Split(s, "\t")
	if col > 0 && col <= len(parts) {
		return parts[col-1]
	}

	return s
}
