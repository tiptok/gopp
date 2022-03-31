package balancer

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log"
	"strings"
	"time"
)

const (
	discoverySchema = "discovery"

	endPointsSplit = ";"

	schemaAddrSplit = "/"
)

var cli *clientv3.Client

type etcdResolver struct {
	rawAddr string
	cc      resolver.ClientConn
}

// NewResolver initialize an etcd client
func NewResolver(etcdAddr string) resolver.Builder {
	return &etcdResolver{rawAddr: etcdAddr}
}

func (r *etcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error

	if err = newClientInstance(r.rawAddr, 15); err != nil {
		return nil, err
	}

	r.cc = cc

	go r.watch("/" + target.Scheme + "/" + target.Authority + "/" + target.Endpoint + "/")
	return r, nil
}

func (r *etcdResolver) Scheme() string {
	return discoverySchema
}

func (r *etcdResolver) watch(keyPrefix string) {
	var addrList []resolver.Address

	getResp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Println(keyPrefix, err)
	}
	for _, v := range getResp.Kvs {
		addrList = append(addrList,
			resolver.Address{Addr: strings.TrimPrefix(string(v.Key), keyPrefix)},
		)
	}

	r.cc.NewAddress(addrList)

	watchChan := cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for item := range watchChan {
		for _, event := range item.Events {
			addr := strings.TrimPrefix(string(event.Kv.Key), keyPrefix)
			switch event.Type {
			case clientv3.EventTypePut:
				if !exist(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					r.cc.NewAddress(addrList)
				}
			case clientv3.EventTypeDelete:
				if newAddrList, ok := remove(addrList, addr); ok {
					addrList = newAddrList
					r.cc.NewAddress(addrList)
				}
			}
		}
	}
}

func (r etcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow") // TODO check
}

// Close closes the resolver.
func (r etcdResolver) Close() {
	log.Println("Close")
}

func exist(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func newClientInstance(etcdAddr string, ttl int64) error {
	var err error
	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(etcdAddr, endPointsSplit),
			DialTimeout: time.Duration(ttl) * time.Second,
		})
		if err != nil {
			return err
		}
	}
	return err
}

// /schema/name/addr
func discoverySchemaAddr(args ...string) string {
	var schemaAddr []string = []string{discoverySchema}
	schemaAddr = append(schemaAddr, args...)
	strBuilder := strings.Builder{}

	strBuilder.WriteString(schemaAddrSplit)
	strBuilder.WriteString(strings.Join(schemaAddr, schemaAddrSplit))
	return strBuilder.String()
}
