package destinations

import (
	"fmt"
	"net"
	"sync"
)

// TCPDestination a tcp destination
type TCPDestination struct {
	Host       string
	Port       string
	m          sync.Mutex
	connection net.Conn
}

// NewTCPDestination initializes the tcp destination
func NewTCPDestination(host, port string) (*TCPDestination, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, err
	}
	return &TCPDestination{host, port, sync.Mutex{}, conn}, nil
}

// Send to the tcp destination
func (t *TCPDestination) Send(msg string) error {
	t.m.Lock()
	_, err := fmt.Fprintf(t.connection, msg+"\n")
	t.m.Unlock()
	return err
}

// Close the tcp destination
func (t *TCPDestination) Close() error {
	return t.connection.Close()
}
