package manager

import "github.com/dexpress/core/module"

// 用于管理挂在在当前Worker上的Task列表，负责加载任务并运行和卸载任务并释放
type WorkerManager struct {
	TaskList []*module.Task
}

// start 运行worker manager
func (wm *WorkerManager) Start() {
}

func (wm *WorkerManager) Stop() {
}

func (wm *WorkerManager) do() {

}

func (wm *WorkerManager) release() {

}
