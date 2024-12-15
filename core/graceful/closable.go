package graceful

type Closable interface {
	Start() error
	Stop() error
}
