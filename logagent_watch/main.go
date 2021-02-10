package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"kenny.com/logagent_watch/conf"
	"kenny.com/logagent_watch/etcd"
	"kenny.com/logagent_watch/kafka"
	"kenny.com/logagent_watch/taillog"
	"kenny.com/logagent_watch/utils"
	"sync"

	"time"
)

var (
	cfg =new(conf.AppConf)
	wg sync.WaitGroup
)
//入口程序

func main()  {
     //1.加载配置文件
	ini.MapTo(cfg,"./conf/config.ini")
	fmt.Println(cfg.KafkaConf.Address)
	//2.初始化kafka连接
	err:=kafka.Init([]string{cfg.KafkaConf.Address},cfg.KafkaConf.ChanMaxSize)

	if err != nil {
		fmt.Println("Init kafka failed, err:\n", err)
		return
	}
    fmt.Println("Init kafka successfully.....")
	//3.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err!=nil{
		fmt.Printf("Init Etcd failed,err:%v\n",err)
		return
	}
	fmt.Println("Init Etcd successfully.....")
	//3.1从etcd中获取日志收集项的配置信息
	//先获取ip，根据ip获取对应的配置，实现每个logagent拉取自己独有的配置
    ipStr,err:=utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey:=fmt.Sprintf(cfg.EtcdConf.Key,ipStr)
	//logEntryConf,err:=etcd.GetConf(cfg.EtcdConf.Key)
	logEntryConf,err:=etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("Get conf failed,err:%v\n",err)
		return
	}
	fmt.Println("Get conf from etcd succeeded,value:")


	for i,v:=range logEntryConf{
		fmt.Printf("index:%v,value:%v\n",i,v)
	}

    //4.收集日志发给kafka
    //4.1遍历每一个收集项，创建tailObj
    taillog.Init(logEntryConf)
	newConfChan:=taillog.NewConfChan()//获取对外可见的通道
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey,newConfChan)//watch发现配置改变会通知上面的通道
	wg.Wait()

}

