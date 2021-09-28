package leas

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func Run(ctx context.Context, endpoints []string) {
	// connect
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalln("Can't create etcd client: ", err)
	}

	// lease
	lease := clientv3.NewLease(client)
	leaseResponse, err := lease.Grant(ctx, 10)
	if err != nil {
		panic(err)
	}
	_, err = client.Put(ctx, "workers/", "online", clientv3.WithLease(leaseResponse.ID))
	if err != nil {
		panic(err)
	}
	leaseKeepAliveResponse, err := lease.KeepAlive(context.Background(), leaseResponse.ID)
	if err != nil {
		panic(err)
	}

	// watch lease
	go func() {
		for keepResp := range leaseKeepAliveResponse {
			fmt.Printf("renew a contract success, Id:%d, TTL:%d, time:%v\n", keepResp.ID, keepResp.TTL, time.Now())
		}
	}()
}
