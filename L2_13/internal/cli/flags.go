package cli

import "flag"

type Options struct {
	Delimiter string // -d
	Fields    string // -f
	Separated bool   // -s
	Filename  string
}

func ParseFlags() Options {
	var opts Options
	flag.StringVar(&opts.Delimiter, "d", "\t", "Delimiter to use for fields")
	flag.StringVar(&opts.Fields, "f", "", "Fields to use for fields")
	flag.BoolVar(&opts.Separated, "s", false, "Show all fields")

	flag.Parse()
	if flag.NArg() > 0 {
		opts.Filename = flag.Arg(0)
	}
	return opts
}
