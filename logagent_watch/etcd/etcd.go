package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//定义一个全局变量
var (
	cli *clientv3.Client

)
//定义一个结构体用于存储获取的配置项
type LogEntry struct {
	Path string`json:"path"`
	Topic string`json:"topic"`

}
//初始化etcd的函数
func Init(addr string,timeout time.Duration)(err error){
	cli, err= clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}

	 return
}

//从etcd中根据key获取配置项
func GetConf(key string)(logEntryConf []*LogEntry,err error)  {
	ctx, cancel:= context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key,clientv3.WithPrefix())

	//resp, err = cli.Get(ctx, "pay")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err=json.Unmarshal(ev.Value,&logEntryConf)
		if err!=nil{
			fmt.Printf("Unmarshal etcd value failed,err:%v\n",err)
			return
		}
	}
	return
}

func WatchConf(key string,newConfCh chan<- []*LogEntry)  {
	watched := cli.Watch(context.Background(), key, clientv3.WithPrefix())
	for wresp:=range watched{
		for _,ev:=range wresp.Events{
			fmt.Printf("type:%v,key:%v,value:%v\n",ev.Type,string(ev.Kv.Key),string(ev.Kv.Value))
		  //通知需要处理方
		  //1.先判断操作的类型
		  var newConf []*LogEntry
			if ev.Type!=clientv3.EventTypeDelete{ //如果不是删除操作，才进行下面的处理
				err:=json.Unmarshal(ev.Kv.Value,&newConf)
				if err != nil {
					fmt.Printf("unmarsharl failed,err:%v",err)
					continue
				}
			}

			fmt.Printf("new conf:%v\n",newConf)
          newConfCh<- newConf
		}
	}
}