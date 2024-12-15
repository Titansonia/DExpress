package manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dexpress/core/module"
	"github.com/dexpress/core/others"
	"go.etcd.io/etcd/clientv3"
)

// 用于管理挂在在当前Worker上的Task列表
type WorkerManager struct {
	Worker   *module.Worker
	TaskList []*module.Task
	TTL      int64 // 租约时间 单位：s

	ctx        context.Context
	cc         chan bool
	wg         *sync.WaitGroup
	etcdClient *clientv3.Client
}

// start 运行worker manager
func (wm *WorkerManager) Start() error {
	wm.wg.Add(1)
	go wm.renewLock()
	return nil
}

func (wm *WorkerManager) Stop() error {
	wm.cc <- true
	fmt.Println("Worker received Stop signal, and waiting for worker to shutdown")
	wm.wg.Wait()
	return nil
}

func (wm *WorkerManager) renewLock() {
	defer wm.wg.Done()
	for {
		select {
		case <-wm.ctx.Done():
			fmt.Println("Worker received shutdown signal")
			return
		case <-wm.cc:
			fmt.Println("Worker received shutdown signal")
			return
		default:
			curWorker := fmt.Sprintf("%s/%d", wm.Worker.IP, wm.Worker.PID)
			others.SetKeyWithTTL(wm.etcdClient, wm.Worker.ID, curWorker, wm.TTL)
			fmt.Printf("Working.... cur worker: %s\n", curWorker)
			// TODO: 这里需要做任务的调度，从任务列表中选择一个任务进行执行

			time.Sleep(5 * time.Second)
		}
	}
}

func NewWorkerManager(ctx context.Context, worker *module.Worker, etcdClient *clientv3.Client) *WorkerManager {
	return &WorkerManager{
		Worker:     worker,
		TaskList:   make([]*module.Task, 0),
		TTL:        30,
		etcdClient: etcdClient,

		ctx: ctx,
		cc:  make(chan bool),
		wg:  &sync.WaitGroup{},
	}
}
