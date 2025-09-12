package commands

import (
	"os"
)

func ChDir(path string) error {
	return os.Chdir(path)
}
