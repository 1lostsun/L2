package cli

import (
	"flag"
	"log"
	"time"
)

type Flags struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func ParseFlags() *Flags {
	host := flag.String("host", "", "server host")
	port := flag.Int("port", 0, "server port")
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")

	flag.Parse()

	if *host == "" || *port == 0 {
		log.Fatalf("host, port and timeout are required")
	}

	return &Flags{
		Host:    *host,
		Port:    *port,
		Timeout: *timeout,
	}
}
