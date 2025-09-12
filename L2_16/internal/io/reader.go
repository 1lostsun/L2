package io

import (
	"bufio"
	"os"
)

func ReadLine() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		return "", scanner.Err()
	}

	return scanner.Text(), scanner.Err()
}
