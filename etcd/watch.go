package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	timeout = time.Second * 10
)

func watcher() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.0.251:49154", "192.168.0.251:12379", "192.168.0.251:49156"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()

	responseCh := cli.Watch(context.Background(), "topic", clientv3.WithPrefix())
	for resp := range responseCh {
		for _, e := range resp.Events {
			fmt.Printf("%s %q: %q\n", e.Type, e.Kv.Key, e.Kv.Value)
		}
	}
}

func put() {
	// expect dial time-out on ipv4 blackhole
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.0.251:49154", "192.168.0.251:12379", "192.168.0.251:49156"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	resp, err := cli.Put(ctx, "topic", "This is a simple test")
	cancel()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
