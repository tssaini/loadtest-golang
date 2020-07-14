package destinations

import (
	"fmt"
	"net"
)

// TCPDestination a tcp destination
type TCPDestination struct {
	Host       string
	Port       string
	connection net.Conn
}

// NewTCPDestination initializes the tcp destination
func NewTCPDestination(host, port string) (*TCPDestination, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, err
	}
	return &TCPDestination{host, port, conn}, nil
}

// Send to the tcp destination
func (t *TCPDestination) Send(msg string) error {
	_, err := fmt.Fprintf(t.connection, msg+"\n")
	return err
}

// Close the tcp destination
func (t *TCPDestination) Close() error {
	return t.connection.Close()
}
