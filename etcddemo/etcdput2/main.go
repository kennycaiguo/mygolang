package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//value:=`[{"path":"c:/tmp/nginx.log","topic":"web_log"},{"path":"d:/dblog/redis.log","topic":"redis_log"},{"path":"d:/dblog/mysql.log","topic":"mysql_log"}]`
	value:=`[{"path":"c:/tmp/nginx.log","topic":"web_log"},{"path":"d:/dblog/redis.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/logagent/192.168.100.11/collect_config", value) //添加前缀是为了方便获取所有key

	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

}
