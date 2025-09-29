package io

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines : Функция чтения строк из STDIN
func ReadLines(filename string) ([]string, error) {
	var file *os.File
	var fileOpenErr error

	if filename != "" {
		if file, fileOpenErr = os.Open(filename); fileOpenErr != nil {
			return nil, fileOpenErr
		}

		defer func(file *os.File) {
			if fileCloseErr := file.Close(); fileCloseErr != nil {
				fmt.Println(fileCloseErr)
				os.Exit(1)
			}
		}(file)
	} else {
		file = os.Stdin
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
