package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"kenny.com/logagent/conf"
	"kenny.com/logagent/etcd"
	"kenny.com/logagent/kafka"
	"kenny.com/logagent/taillog"
	"time"
)

var (
	cfg =new(conf.AppConf)
)
//入口程序

func main()  {
     //1.加载配置文件
	ini.MapTo(cfg,"./conf/config.ini")
	fmt.Println(cfg.KafkaConf.Address)
	//2.初始化kafka连接
	err:=kafka.Init([]string{cfg.KafkaConf.Address},cfg.KafkaConf.ChanMaxSize)

	if err != nil {
		fmt.Println("Init kafka failed, err:", err)
		return
	}
    fmt.Println("Init kafka successfully.....")
	//3.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err!=nil{
		fmt.Printf("Init Etcd failed,err:%v",err)
		return
	}
	fmt.Println("Init Etcd successfully.....")
	//3.1从etcd中获取日志收集项的配置信息
	logEntryConf,err:=etcd.GetConf("/xxx")
	if err != nil {
		fmt.Printf("Get conf failed,err:%v\n",err)
		return
	}
	fmt.Println("Get conf from etcd succeeded,value:")
	for i,v:=range logEntryConf{
		fmt.Printf("index:%v,value:%v\n",i,v)
	}
	//3.2 watch日志收集项的变化，实现热加载

    //4.收集日志发给kafka
    //4.1遍历每一个收集项，创建tailObj
    taillog.Init(logEntryConf)

	////3.打开日志文件准备收集
	//err = taillog.Init(cfg.TaillogConf.FileName)
	//if err!=nil{
	//	fmt.Println("Init taillog failed, err:", err)
	//	return
	//}
	//fmt.Println("Init taillog successfully.....")
    //4.业务逻辑
	//run()
}

//func run() {
//
//	//1.读取日志
//
//	for {
//		select {
//		case line :=<- taillog.ReadChan():
//			//2.发送到kafka
//			kafka.SendToKafka(cfg.KafkaConf.Topic,line.Text)
//		default:
//			time.Sleep(time.Second)
//	   }
//	}
//
//}
