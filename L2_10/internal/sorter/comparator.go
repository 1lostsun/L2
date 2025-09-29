package sorter

import (
	"github.com/1lostsun/L2/tree/main/L2_10/internal/cli"
	"strconv"
	"strings"
)

var monthMap = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
	"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
	"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
}

func getKey(s string, col int, ignoreTrails bool) string {
	if ignoreTrails {
		s = strings.TrimRight(s, "\t")
	}

	parts := strings.Split(s, "\t")
	if col > 0 && col <= len(parts) {
		return parts[col-1]
	}

	if len(parts) > 0 {
		return parts[0]
	}

	return s
}

// ParseHuman преобразует строку с суффиксами размеров (K, M, G) в число байт.
func ParseHuman(s string) (int64, bool) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0, false
	}

	mult := int64(1)
	last := s[len(s)-1]
	num := s[:len(s)-1]

	switch last {
	case 'K', 'k':
		mult = 1024
	case 'M', 'm':
		mult = 1024 * 1024
	case 'G', 'g':
		mult = 1024 * 1024 * 1024
	}

	val, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0, false
	}

	return val * mult, true
}

func makeComparator(opt cli.Options) func(a, b string) bool {
	return func(a, b string) bool {
		ka := getKey(a, opt.Column, opt.IgnoreTrails)
		kb := getKey(b, opt.Column, opt.IgnoreTrails)

		if opt.Human {
			ha, okA := ParseHuman(ka)
			hb, okB := ParseHuman(kb)

			if okA && okB {
				if opt.Reverse {
					return ha > hb
				}
				return ha < hb
			}

			if okA && !okB {
				return !opt.Reverse
			}
			if !okA && okB {
				return opt.Reverse
			}
		}

		if opt.Numeric {
			na, errA := strconv.Atoi(ka)
			nb, errB := strconv.Atoi(kb)
			if errA == nil && errB == nil {
				if opt.Reverse {
					return na > nb
				}
				return na < nb
			}
		}

		if opt.Months {
			mA, okA := monthMap[ka]
			mB, okB := monthMap[kb]

			if okA && okB {
				if opt.Reverse {
					return mA > mB
				}
				return mA < mB
			}

			if okA && !okB {
				return !opt.Reverse
			}
			if !okA && okB {
				return opt.Reverse
			}
		}

		if opt.Reverse {
			return ka > kb
		}
		return ka < kb
	}
}
