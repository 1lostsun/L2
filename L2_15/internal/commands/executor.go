package commands

import (
	"fmt"
	"github.com/gammazero/deque"
	"os/exec"
	"strings"
)

func Executor(deq *deque.Deque[[]string]) {
	deq.IterPopFront()(func(args []string) bool {
		cmd := strings.ToLower(args[0])
		switch cmd {
		case "cd":
			if err := ChDir(args[1]); err != nil {
				return false
			}
			fmt.Println("executing cd command", args[1:])
		case "pwd":
			currentDir, err := Pwd()
			if err != nil {
				return false
			}
			fmt.Println("executing pwd command")
			fmt.Println("current directory: ", currentDir)
		case "echo":
			Echo(args[1:])
			fmt.Println("executing echo command")
		case "kill":
			if err := Kill(args[1]); err != nil {
				return false
			}
			fmt.Println("executing kill command", args[1:])
		case "ps":
			if psErr := Ps(args[1:]); psErr != nil {
				return false
			}
			fmt.Println("executing ps command", args[1:])
		default:
			out, execErr := exec.Command(args[0], args[1:]...).CombinedOutput()
			if execErr != nil {
				return false
			}

			fmt.Println(string(out))
		}
		return true
	})
}
