package commands

import (
	"fmt"
	"strconv"
	"syscall"
)

func Kill(PID string) error {
	pid, pidConvErr := strconv.Atoi(PID)
	if pidConvErr != nil {
		return pidConvErr
	}
	if pidKillErr := syscall.Kill(pid, syscall.SIGTERM); pidKillErr != nil {
		return pidKillErr
	}

	fmt.Println("pid:", pid, " was killed")
	return nil
}
