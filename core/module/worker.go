package module

// 一个DTS进程，同一个IP上支持部署多个DTS进程
type Worker struct {
	ID    string // worker id
	IP    string // worker ip
	PID   int    // worker pid
	Group string // group worker belongs to
}
