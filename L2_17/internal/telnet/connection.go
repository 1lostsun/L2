package telnet

import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_17/internal/cli"
	"net"
)

// Connect : Функция установки соединения
func Connect(flags *cli.Flags) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%d", flags.Host, flags.Port)
	conn, err := net.DialTimeout("tcp", addr, flags.Timeout)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
