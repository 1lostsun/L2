package commands

import (
	"fmt"
	"os"
)

// Pwd : Функция возвращает текущую директорию
func Pwd() (string, error) {
	currentDir, CurrentDirErr := os.Getwd()
	if CurrentDirErr != nil {
		return "", fmt.Errorf("could not get current directory: %w", CurrentDirErr)
	}

	return currentDir, nil
}
