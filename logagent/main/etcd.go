package main

import (
    "github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
)

func initEtcd(addr string) (err error) {
	cli, err := etcd_client.New(clientv3.Config{
		Endpoints: []string{"localhost:2379","localhost:22379","localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed ,err",err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()
}
