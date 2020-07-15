package destinations

import (
	"fmt"
	"net"
	"sync"
)

// UDPDestination a udp destination
type UDPDestination struct {
	Host       string
	Port       string
	m          sync.Mutex
	connection net.Conn
}

// NewUDPDestination initializes the udp destination
func NewUDPDestination(host, port string) (*UDPDestination, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, err
	}
	return &UDPDestination{host, port, sync.Mutex{}, conn}, nil
}

// Send to the tcp destination
func (t *UDPDestination) Send(msg string) error {
	t.m.Lock()
	_, err := fmt.Fprintf(t.connection, msg+"\n")
	t.m.Unlock()
	return err
}

// Close the tcp destination
func (t *UDPDestination) Close() error {
	return t.connection.Close()
}
