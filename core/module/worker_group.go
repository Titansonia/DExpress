package module

// 管理Worker，每个Group下包含若干个Worker，实现Worker级的高可用
type WorkerGroup struct {
	ID        string
	WorkerMap map[string]bool //worker id -> isActive
}
