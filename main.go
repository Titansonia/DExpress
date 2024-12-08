package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/dexpress/core/others"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	// fmt.Println("hello")
	// ctx, cancel := context.WithCancel(context.Background())
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go worker(ctx, &wg)

	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, os.Interrupt)

	// <-stop
	// fmt.Println("Shuttting down...")

	// cancel()
	// wg.Wait()
	// fmt.Println("Shutdown Complete")
	callEtcd()
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker received shutdown signal")
			return
		default:
			// 执行工作任务
			fmt.Println("Working....")
			time.Sleep(time.Second)
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
