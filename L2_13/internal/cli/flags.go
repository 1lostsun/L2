package cli

import (
	"flag"
	"strconv"
	"strings"
)

// Options : Структура флагов
type Options struct {
	Delimiter string // -d
	Fields    []int  // -f
	Separated bool   // -s
	Filename  string
}

// ParseFlags : Функция парсинга флагов
func ParseFlags() Options {
	var delimiter, fields string
	var separated bool

	flag.StringVar(&delimiter, "d", "\t", "Delimiter to use for fields")
	flag.StringVar(&fields, "f", "", "Fields to use for fields")
	flag.BoolVar(&separated, "s", false, "Show all fields")

	flag.Parse()
	opts := Options{
		Delimiter: delimiter,
		Fields:    ParseList(fields),
		Separated: separated,
	}
	return opts
}

// ParseList : преобразует строку с числами и диапазонами в список целых чисел
func ParseList(s string) []int {
	var res []int
	for _, part := range strings.Split(s, ",") {
		if strings.TrimSpace(part) == "" {
			continue
		}
		if strings.Contains(part, "-") {
			splitPart := strings.Split(part, "-")
			start, err1 := strconv.Atoi(splitPart[0])
			end, err2 := strconv.Atoi(splitPart[1])
			if err1 == nil && err2 == nil && start <= end {
				for i := start; i <= end; i++ {
					res = append(res, i)
				}
			}
		} else {
			if n, err := strconv.Atoi(part); err == nil {
				res = append(res, n)
			}
		}
	}
	return res
}
