package manager

import (
	"sync"
	"time"

	"github.com/dexpress/core/module"
)

const (
	gapToReloadTask = 30 * time.Second
)

// 负责管理一个具体Task的运行状态
type TaskManager struct {
	group   string
	dbName  string
	shardNo string

	// 缓存一份任务信息
	task    *module.Task
	sigQuit chan bool
	wg      *sync.WaitGroup
}

// 外部调用，用于当前goroutine的优雅推出
func (tm *TaskManager) Stop() {
	tm.sigQuit <- true
	tm.wg.Wait()
}

// 外部调用，用于启动当前goroutine
func (tm *TaskManager) Start() {
	tm.wg.Add(1)
	// 管理任务运行时状态，启动或暂停
	go tm.manage()

}

func (tm *TaskManager) manage() {
	defer tm.wg.Done()
	for {
		select {
		case <-time.After(gapToReloadTask):
			tm.reload()
		case <-tm.sigQuit:
			tm.release()
		}

	}
}

// reload 重新从存储中加载Task元信息
func (tm *TaskManager) reload() error {
	// 重新加载任务元信息，并启动任务
	return nil
}

// 当前goutine调用，用于释放依赖资源
func (tm *TaskManager) release() {

}
