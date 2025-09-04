package cli

import "flag"

type Options struct {
	Column       int
	Numeric      bool
	Reverse      bool
	Unique       bool
	IgnoreTrails bool
	Filename     string
}

func ParseFlags() Options {
	k := flag.Int("k", 0, "Number of columns to generate")
	n := flag.Bool("n", false, "Number of rows to generate")
	r := flag.Bool("r", false, "Reverse order")
	u := flag.Bool("u", false, "Unique order")
	b := flag.Bool("b", false, "Ignore trails")

	flag.Parse()
	var filename string
	if flag.NArg() > 0 {
		filename = flag.Arg(0)
	}

	return Options{
		Column:       *k,
		Numeric:      *n,
		Reverse:      *r,
		Unique:       *u,
		IgnoreTrails: *b,
		Filename:     filename,
	}
}
