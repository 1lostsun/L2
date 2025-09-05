package filter

import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_12/internal/cli"
	"regexp"
	"strconv"
)

func Apply(lines []string, opt cli.Options) []string {
	var result []string
	toPrint := make(map[int]bool)
	re := compileRegex(opt)
	for i, line := range lines {
		match := re.MatchString(line)
		if opt.Invert {
			match = !match
		}

		if match {
			toPrint[i] = true
			before := opt.Before

			if opt.Cross > before {
				before = opt.Cross
			}

			for j := 1; j <= before; j++ {
				if i-j >= 0 {
					toPrint[i-j] = true
				}
			}

			after := opt.After
			if opt.Cross > after {
				after = opt.Cross
			}

			for j := 1; j <= after; j++ {
				if i+j <= len(lines) {
					toPrint[i+j] = true
				}
			}
		}
	}

	for i, line := range lines {
		if toPrint[i] {
			if opt.LineNumber {
				line = fmt.Sprintf("%d:%s", i+1, line)
			}
			result = append(result, line)
		}
	}

	if opt.Count {
		return []string{strconv.Itoa(len(result))}
	}

	return result
}

func compileRegex(opt cli.Options) *regexp.Regexp {
	pattern := opt.Pattern

	if opt.Fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	if opt.IgnoreCase {
		pattern = `(?i)` + pattern
	}

	re := regexp.MustCompile(pattern)

	return re
}
