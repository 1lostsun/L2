package io

import (
	"bufio"
	"fmt"
	"net"
)

func Reader(conn net.Conn, done chan struct{}) error {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	close(done)
	return scanner.Err()
}
