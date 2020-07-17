package destinations

import (
	"fmt"
	"net"
	"sync"
)

// TCPConn a tcp connection
type TCPConn struct {
	Host       string
	Port       string
	m          sync.Mutex
	connection net.Conn
}

// NewTCPConn initializes a tcp connection
func NewTCPConn(host, port string) (*TCPConn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, err
	}
	return &TCPConn{host, port, sync.Mutex{}, conn}, nil
}

// Send to the tcp destination
func (t *TCPConn) Send(msg string) error {
	t.m.Lock()
	_, err := fmt.Fprintf(t.connection, msg+"\n")
	t.m.Unlock()
	return err
}

// Close the tcp destination
func (t *TCPConn) Close() error {
	return t.connection.Close()
}
