package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/dexpress/core/manager"
	"github.com/dexpress/core/module"
	"github.com/dexpress/core/others"
	"go.etcd.io/etcd/clientv3"
)

func main() {

	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd 服务地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}
	defer cli.Close()
	fmt.Println("hello")
	ctx, cancel := context.WithCancel(context.Background())
	WorkerManager := manager.NewWorkerManager(ctx, &module.Worker{
		ID:    "myworker",
		IP:    "",
		PID:   os.Getpid(),
		Group: "",
	}, cli)
	fmt.Println("main send start signal.")
	go WorkerManager.Start()
	time.Sleep(10 * time.Second)
	fmt.Println("main send stop signal.")
	go WorkerManager.Stop()
	time.Sleep(10 * time.Second)
	fmt.Println("main send start signal.")
	go WorkerManager.Start()
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(10 * time.Second)
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go worker(ctx, &wg, cli)

	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, os.Interrupt)

	// <-stop
	// fmt.Println("Shuttting down...")

	// cancel()
	// wg.Wait()
	// fmt.Println("Shutdown Complete")
	// // callEtcd()
}

func worker(ctx context.Context, wg *sync.WaitGroup, cli *clientv3.Client) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker received shutdown signal")
			return
		default:
			// 执行工作任务
			pid := os.Getpid()
			value, _ := others.GetKey(cli, "myLock")
			if value == "" || value == strconv.Itoa(pid) {
				others.SetKeyWithTTL(cli, "myLock", strconv.Itoa(pid), 30)
				fmt.Println(fmt.Sprintf("Working.... cur PID: %d", pid))
			} else {
				oldPid, _ := strconv.Atoi(value)
				if oldPid != pid {
					fmt.Println(fmt.Sprintf("Other process is holding the lock, PID: %d", oldPid))
				}
			}

			time.Sleep(40 * time.Second)
		}
	}

}

func callEtcd() {
	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd 服务地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}
	defer cli.Close()

	// 设置键值对并指定过期时间
	others.SetKeyWithTTL(cli, "mykey", "myvalue", 10)

	// 获取键值对
	others.GetKey(cli, "mykey")

	// 更新键值对
	others.UpdateKey(cli, "mykey", "newvalue")

	// 删除键值对
	others.DeleteKey(cli, "mykey")
}
