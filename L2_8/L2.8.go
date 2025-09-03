package main

import (
	"github.com/1lostsun/L2/L2.8/ntp_utils"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	ntpTime, err := ntp_utils.Now()
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Fatal(err)
	}

	log.Println("Точное время:", ntpTime)
}
