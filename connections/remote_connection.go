package connections

// RemoteConn all remote connections should implement the Send() and Close() functions
type RemoteConn interface {
	Send(msg string) error
	Close() error
}
