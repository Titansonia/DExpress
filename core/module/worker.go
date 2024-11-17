package module

// 一个DTS进程，同一个IP上支持部署多个DTS进程
type Worker struct {
	ID string
	IP string
}
