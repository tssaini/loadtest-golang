package main

import (
	"log"

	"github.com/tssaini/syslog-ng-config-testing/destinations"
)

func main() {
	tcpConn, err := destinations.NewTCPConn("127.0.0.1", "601")
	if err != nil {
		log.Fatalf("Unable to create udp destiantion")
	}
	defer tcpConn.Close()

	// err = tcpDest.Send("Hello world!")
	// err = tcpDest.Send("udp works!")
	// err = tcpDest.Send("cool it works again and again")
	// if err != nil {
	// 	log.Fatalf("Unable to send to udp destnation")
	// }
	err = generate("Hello World!", 200, tcpConn)
	if err != nil {
		log.Fatalf("unable to generate logs %v", err)
	}
}
