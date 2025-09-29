package commands

import (
	"fmt"
	"strings"
)

// Echo : Функция для вывода строк в консоль эквивалетная линукс команде echo
func Echo(args []string) {
	newline := true
	escape := false
	var out []string

	for _, a := range args {
		if strings.HasPrefix(a, "-") && len(out) == 0 {
			switch a {
			case "-n":
				newline = false
				continue
			case "-e":
				escape = true
				continue
			}
		}
		out = append(out, a)
	}

	text := strings.Join(out, " ")

	if escape {
		text = strings.ReplaceAll(text, `\n`, "\n")
		text = strings.ReplaceAll(text, `\t`, "\t")
	}

	if newline {
		fmt.Println(text)
	} else {
		fmt.Print(text)
	}
}
