package main

import (
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
