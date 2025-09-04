package main

import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_10/internal/cli"
	"github.com/1lostsun/L2/tree/main/L2_10/internal/io"
	"github.com/1lostsun/L2/tree/main/L2_10/internal/sorter"
	"log"
	"os"
)

func main() {
	opts := cli.ParseFlags()

	in, err := io.OpenInput(opts.Filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(in)

	lines, err := io.ReadLines(in)
	if err != nil {
		log.Fatal(err)
	}

	if opts.Checked {
		if !sorter.IsSorted(lines, opts) {
			log.Fatalf("date isn't sorted")
		}

		fmt.Println("date sorted")
	}

	lines = sorter.SortLines(lines, opts)
	for _, l := range lines {
		fmt.Println(l)
	}
}
