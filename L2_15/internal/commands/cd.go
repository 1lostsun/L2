package commands

import (
	"os"
)

// ChDir : Функция смены директории
func ChDir(path string) error {
	return os.Chdir(path)
}
