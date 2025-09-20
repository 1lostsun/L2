package main

import (
	"github.com/1lostsun/L2/tree/main/L2_17/internal/cli"
	"github.com/1lostsun/L2/tree/main/L2_17/internal/io"
	"github.com/1lostsun/L2/tree/main/L2_17/internal/telnet"
	"log"
)

func main() {
	done := make(chan struct{})
	flags := cli.ParseFlags()
	connection, connErr := telnet.Connect(flags)
	if connErr != nil {
		log.Fatal(connErr)
	}

	go func() {
		err := io.Writer(connection, done)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := io.Reader(connection, done)
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-done
}
