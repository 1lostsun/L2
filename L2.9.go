package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	tests := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
	}

	for _, t := range tests {
		res, err := unpackStringExtended(t)
		if err != nil {
			fmt.Printf("'%s' -> ошибка: %v\n", t, err)
		} else {
			fmt.Printf("'%s' -> '%s'\n", t, res)
		}
	}
}

// 1 вариант без доп задания
func unpackString(s string) (string, error) {
	newStr := ""
	for i := 0; i < len(s); i++ {
		symbol := s[i]
		if unicode.IsDigit(rune(symbol)) {
			if i > 0 && unicode.IsLetter(rune(s[i-1])) {
				n, err := strconv.Atoi(string(symbol))
				if err != nil {
					return "", err
				}
				newStr += strings.Repeat(string(s[i-1]), n)
			} else {
				return "", fmt.Errorf("invalid string")
			}
		} else {
			newStr += string(symbol)
		}
	}
	return newStr, nil
}

// 2 вариант с доп заданием
func unpackStringExtended(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var newStr strings.Builder
	runes := []rune(s)
	escaped := false
	var lastRune rune
	lastCharWritten := false

	for i := 0; i < len(runes); i++ {
		current := runes[i]

		if escaped {
			newStr.WriteRune(current)
			escaped = false
			lastRune = current
			lastCharWritten = true
			continue
		}

		if current == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(current) {
			if !lastCharWritten {
				return "", fmt.Errorf("invalid string: string starts with digital")
			}
			n, err := strconv.Atoi(string(current))
			if err != nil {
				return "", err
			}
			newStr.WriteString(strings.Repeat(string(lastRune), n-1))
			continue
		}

		newStr.WriteRune(current)
		lastRune = current
		lastCharWritten = true
	}

	if escaped {
		return "", fmt.Errorf("invalid string: ends with escape character")
	}

	return newStr.String(), nil
}
