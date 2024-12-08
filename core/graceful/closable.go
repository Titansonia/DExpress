package graceful

type Closable interface {
	Start() error
	Close() error
}
