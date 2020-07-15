package main

import (
	"log"

	"github.com/tssaini/syslog-ng-config-testing/destinations"
)

func main() {
	tcpDest, err := destinations.NewTCPDestination("127.0.0.1", "601")
	if err != nil {
		log.Fatalf("Unable to create udp destiantion")
	}
	defer tcpDest.Close()
	// err = tcpDest.Send("Hello world!")
	// err = tcpDest.Send("udp works!")
	// err = tcpDest.Send("cool it works again and again")
	// if err != nil {
	// 	log.Fatalf("Unable to send to udp destnation")
	// }
	err = generate("Hello World!", 200, tcpDest)
	if err != nil {
		log.Fatalf("unable to generate logs %v", err)
	}
}
