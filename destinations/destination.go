package destinations

// Destination all destination should implement the Send() and Close() functions
type Destination interface {
	Send(msg string) error
	Close() error
}
