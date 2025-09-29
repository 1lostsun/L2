package cut

import (
	"github.com/1lostsun/L2/tree/main/L2_13/internal/cli"
	"strings"
)

// Cut : выбирает указанные поля из строки по разделителю.
func Cut(line string, opts cli.Options) string {
	if opts.Separated && !strings.Contains(line, opts.Delimiter) {
		return ""
	}

	splitLine := strings.Split(line, opts.Delimiter)
	var res []string

	for _, field := range opts.Fields {
		if field-1 < len(splitLine) {
			res = append(res, splitLine[field-1])
		}
	}

	return strings.Join(res, opts.Delimiter)
}
