package etcd

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/atomic"
)

func RegisterService(context context.Context, etcdaddrs []string, serviceName string) error {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdaddrs,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}

	kv := clientv3.NewKV(c)
	lease := clientv3.NewLease(c)
	var curLeaseId clientv3.LeaseID = 0
	timer := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-timer.C:
				if curLeaseId == 0 {
					leaseResp, err := lease.Grant(context, 10)
					if err != nil {
						panic(err)
					}

					key := "/services/" + serviceName + "/" + fmt.Sprintf("%d", leaseResp.ID)
					if _, err := kv.Put(context, key, value, clientv3.WithLease(leaseResp.ID)); err != nil {
						panic(err)
					}
					curLeaseId = leaseResp.ID
				} else {
					// 续约租约，如果租约已经过期将curLeaseId复位到0重新走创建租约的逻辑
					if _, err := lease.KeepAliveOnce(context, curLeaseId); err == rpctypes.ErrLeaseNotFound {
						curLeaseId = 0
						continue
					}
				}
			case <-context.Done():
				return
			}
		}
		defer c.Close()
	}()

	return nil
}

func GetService(context context.Context, etcdaddrs []string, serviceName string) *remoteService {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdaddrs,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	kv := clientv3.NewKV(c)
	rangeResp, err := kv.Get(context, "/services/"+serviceName, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	service.mutex.Lock()
	for _, kv := range rangeResp.Kvs {
		service.nodes[string(kv.Key)] = string(kv.Value)
	}
	service.mutex.Unlock()
	watcher := clientv3.NewWatcher(client)
	watchChan := watcher.Watch(context, service.name, clientv3.WithPrefix())

	go func() {
		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				service.mutex.Lock()
				switch event.Type {
				case mvccpb.PUT: //PUT事件，目录下有了新key
					service.nodes[string(event.Kv.Key)] = string(event.Kv.Value)
				case mvccpb.DELETE: //DELETE事件，目录中有key被删掉(Lease过期，key 也会被删掉)
					delete(service.nodes, string(event.Kv.Key))
				}
				service.mutex.Unlock()
			}
		}
	}()
}

type RoundRipper struct {
	host  string
	index atomic.Int64
	ips   []string
}

func (r *RoundRipper) GetIP() string {
	r1 := r.index.Add(1)
	return r.ips[r1%len(r.ips)]
}

func NewRoundRipper(host string, ips []string) *RoundRipper {
	return &RoundRipper{index: 0, ips: ips}
}
