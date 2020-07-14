package destinations

type Destination interface {
	Send(msg string) error
	Close() error
}