package cli

import "flag"

// Options : Структура флагов и настроек
type Options struct {
	Column       int
	Numeric      bool
	Reverse      bool
	Unique       bool
	IgnoreTrails bool
	Months       bool
	Checked      bool
	Human        bool
	Filename     string
}

// ParseFlags : Функция парсинга флагов
func ParseFlags() Options {
	k := flag.Int("k", 0, "Number of columns to generate")
	n := flag.Bool("n", false, "Number of rows to generate")
	r := flag.Bool("r", false, "Reverse order")
	u := flag.Bool("u", false, "Unique order")
	b := flag.Bool("b", false, "Ignore trails")
	M := flag.Bool("M", false, "Months order")
	c := flag.Bool("c", false, "Checked order")
	h := flag.Bool("h", false, "Human readable order")

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
		Months:       *M,
		Checked:      *c,
		Human:        *h,
		Filename:     filename,
	}
}
