package cli

import (
	"flag"
	"fmt"
	"os"
)

// Options : Структура флагов и настроек
type Options struct {
	Pattern    string
	Filename   string
	After      int  // -A N
	Before     int  // -B N
	Cross      int  // -C N == -A N -B N
	Count      bool // -c
	IgnoreCase bool // -i
	Invert     bool // -v
	Fixed      bool // -F
	LineNumber bool // -n
}

// ParseFlags : Функция парсинга флагов
func ParseFlags() Options {
	var opt Options
	flag.IntVar(&opt.After, "A", 0, "Number of lines to print")
	flag.IntVar(&opt.Before, "B", 0, "Number of lines to print")
	flag.IntVar(&opt.Cross, "C", 0, "Number of lines to print")
	flag.BoolVar(&opt.Count, "c", false, "Number of lines to print")
	flag.BoolVar(&opt.IgnoreCase, "i", false, "Ignore case sensitivity")
	flag.BoolVar(&opt.Invert, "v", false, "Invert")
	flag.BoolVar(&opt.Fixed, "F", false, "Fixed")
	flag.BoolVar(&opt.LineNumber, "n", false, "Line number")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: l2_12 [options] pattern [filename]")
		os.Exit(1)
	}

	opt.Pattern = args[0]

	if len(args) > 1 {
		opt.Filename = args[1]
	}

	return opt
}
