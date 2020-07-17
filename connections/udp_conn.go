package connections

import (
	"fmt"
	"net"
	"sync"
)

// UDPConn a udp connections
type UDPConn struct {
	Host       string
	Port       string
	m          sync.Mutex
	connection net.Conn
}

// NewUDPConn initializes the udp connection
func NewUDPConn(host, port string) (*UDPConn, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, err
	}
	return &UDPConn{host, port, sync.Mutex{}, conn}, nil
}

// Send to the udp connection
func (t *UDPConn) Send(msg string) error {
	t.m.Lock()
	_, err := fmt.Fprintf(t.connection, msg+"\n")
	t.m.Unlock()
	return err
}

// Close the udp connection
func (t *UDPConn) Close() error {
	return t.connection.Close()
}
