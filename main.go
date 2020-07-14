package main

import (
	"log"

	"github.com/tssaini/syslog-ng-config-testing/destinations"
)

func main() {
	tcpDest, err := destinations.NewTCPDestination("127.0.0.1", "8080")
	if err != nil {
		log.Fatalf("Unable to create tcp destiantion")
	}
	defer tcpDest.Close()
	err = tcpDest.Send("Hello world!")
	if err != nil {
		log.Fatalf("Unable to send to tcp destnation")
	}
}
