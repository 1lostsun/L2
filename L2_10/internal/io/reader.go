package io

import (
	"bufio"
	"os"
)

// OpenInput : Функция открытия переданного файла для чтения
func OpenInput(filename string) (*os.File, error) {
	if filename != "" {
		return os.Open(filename)
	}

	return os.Stdin, nil
}

// ReadLines : Чтение построчно из файла
func ReadLines(f *os.File) ([]string, error) {
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
