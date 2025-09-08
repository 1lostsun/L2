package main

import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_13/internal/cli"
	"github.com/1lostsun/L2/tree/main/L2_13/internal/cut"
	"github.com/1lostsun/L2/tree/main/L2_13/internal/io"
	"log"
)

func main() {
	opts := cli.ParseFlags()
	lines, err := io.ReadLines(opts.Filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		cutLine := cut.Cut(line, opts)
		if len(cutLine) > 0 {
			fmt.Println(cutLine)
		}
	}
}
