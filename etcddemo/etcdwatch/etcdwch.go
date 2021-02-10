package main
//Watch实例程序。watch只对更改有监视，对查询没有反应
import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err!=nil{
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd successfully...")
	defer client.Close()
	watched := client.Watch(context.Background(), "/job/", clientv3.WithPrefix())
	for wresp:=range watched{
		for _,ev:=range wresp.Events{
           fmt.Printf("type:%v,key:%v,value:%v\n",ev.Type,string(ev.Kv.Key),string(ev.Kv.Value))
		}
	}
}
