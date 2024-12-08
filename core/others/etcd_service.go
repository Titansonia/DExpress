package others

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// 设置键值对并指定过期时间
func SetKeyWithTTL(cli *clientv3.Client, key, value string, ttl int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 创建一个租约
	leaseResp, err := cli.Grant(ctx, ttl)
	if err != nil {
		log.Fatalf("failed to grant lease: %v", err)
	}

	// 设置键值对
	_, err = cli.Put(ctx, key, value, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		log.Fatalf("failed to put key: %v", err)
	}

	fmt.Printf("Set key %s with value %s and TTL %d seconds\n", key, value, ttl)
}

// 获取键值对
func GetKey(cli *clientv3.Client, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := cli.Get(ctx, key)
	if err != nil {
		log.Fatalf("failed to get key: %v", err)
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("Key: %s, Value: %s\n", ev.Key, ev.Value)
	}
}

// 更新键值对
func UpdateKey(cli *clientv3.Client, key, newValue string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cli.Put(ctx, key, newValue)
	if err != nil {
		log.Fatalf("failed to update key: %v", err)
	}

	fmt.Printf("Updated key %s to new value %s\n", key, newValue)
}

// 删除键值对
func DeleteKey(cli *clientv3.Client, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cli.Delete(ctx, key)
	if err != nil {
		log.Fatalf("failed to delete key: %v", err)
	}

	fmt.Printf("Deleted key %s\n", key)
}
