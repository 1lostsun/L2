package io

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	if filename == "" {
		return readFromStdin()
	}

	return readFromFile(filename)
}

func readFromFile(filename string) ([]string, error) {
	file, fileOpenErr := os.Open(filename)
	if fileOpenErr != nil {
		return nil, fmt.Errorf("error opening file: %v", fileOpenErr)
	}

	defer func(file *os.File) {
		fileCloseErr := file.Close()
		if fileCloseErr != nil {
			fmt.Println(fileCloseErr)
			os.Exit(1)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}

	return lines, nil
}

func readFromStdin() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning stdin: %w", err)
	}

	return lines, nil
}
