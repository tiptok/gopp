package balancer

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func Register(etcdAddr, name, addr string, ttl int64) error {
	var err error
	if err = newClientInstance(etcdAddr, 15); err != nil {
		return err
	}

	ticker := time.NewTimer(time.Second * time.Duration(ttl))

	go func() {
		for {
			getResp, err := cli.Get(context.Background(), discoverySchemaAddr(name, addr))
			if err != nil {
				log.Println(err)
			} else if getResp.Count == 0 {
				err = withAlive(name, addr, ttl)
				if err != nil {
					log.Println(name, addr, err)
				}
			}
			<-ticker.C
		}
	}()

	return err
}

func withAlive(name, addr string, ttl int64) error {
	leaseResp, err := cli.Grant(context.Background(), ttl)

	if err != nil {
		return err
	}
	fmt.Println("register key:", discoverySchemaAddr(name, addr))

	_, err = cli.Put(context.Background(), discoverySchemaAddr(name, addr), addr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}

	_, err = cli.KeepAlive(context.Background(), leaseResp.ID)
	return err
}

func UnRegister(name, addr string) {
	if cli != nil {
		//cli.Revoke(context.Background())
		fmt.Println("unregister key:", discoverySchemaAddr(name, addr))
		cli.Delete(context.Background(), discoverySchemaAddr(name, addr))
	}
}
