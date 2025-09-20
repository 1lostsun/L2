package io

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Writer(conn net.Conn, done chan struct{}) error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Fprintln(conn, line)
		if err != nil {
			return err
		}
	}
	close(done)
	conn.Close()

	return scanner.Err()
}
