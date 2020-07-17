package main

import (
	"log"
)

func main() {
	host := "127.0.0.1"
	port := "601"
	conns, err := createConns(host, port, "tcp", 10)
	if err != nil {
		log.Fatalf("Unable to create connections")
	}
	generateLogs(conns, 20, "Hello google world!")
}
