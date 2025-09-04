package sorter

import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_10/internal/cli"
	"sort"
)

func SortLines(lines []string, opt cli.Options) []string {
	cmp := makeComparator(opt)
	sort.SliceStable(lines, func(i, j int) bool {
		return cmp(lines[i], lines[j])
	})

	if opt.Unique {
		var uniq []string
		for i, line := range lines {
			if i == 0 || line != lines[i-1] {
				uniq = append(uniq, line)
			}
		}
		lines = uniq
	}

	return lines
}

func IsSorted(lines []string, opt cli.Options) bool {
	cmp := makeComparator(opt)
	for i := 1; i < len(lines); i++ {
		if cmp(lines[i], lines[i-1]) {
			return false
		}
	}
	return true
}

func Debug(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
