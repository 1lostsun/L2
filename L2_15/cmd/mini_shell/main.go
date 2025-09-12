package main

import (
	"github.com/1lostsun/L2/tree/main/L2_15/internal/commands"
	"github.com/1lostsun/L2/tree/main/L2_15/internal/io"
	"github.com/fatih/color"
	"log"
	"os"
)

func main() {
	for {
		currentDir, CurrentDirErr := os.Getwd()
		if CurrentDirErr != nil {
			log.Fatal(CurrentDirErr)
		}
		dirColor := color.New(color.FgCyan)
		_, dirPrintErr := dirColor.Print(currentDir, " $")
		if dirPrintErr != nil {
			log.Fatal(dirPrintErr)
		}

		input, scannerErr := io.Reader()
		if scannerErr != nil {
			log.Fatal(scannerErr)
		}
		commands.Executor(input)
	}
}
