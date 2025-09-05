package main

import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_12/internal/cli"
	"github.com/1lostsun/L2/tree/main/L2_12/internal/filter"
	"github.com/1lostsun/L2/tree/main/L2_12/internal/io"
	"log"
)

func main() {
	opts := cli.ParseFlags()
	lines, err := io.ReadLines(opts.Filename)
	if err != nil {
		log.Fatal(err)
	}

	results := filter.Apply(lines, opts)
	for _, result := range results {
		fmt.Println(result)
	}
}
