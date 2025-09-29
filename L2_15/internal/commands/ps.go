package commands

import (
	"fmt"
	"os/exec"
)

// Ps : Функция выводит список запущенных процессов
func Ps(args []string) error {
	out, err := exec.Command("ps", args...).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
