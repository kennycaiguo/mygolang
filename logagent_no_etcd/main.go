package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"kenny.com/logagent/conf"
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
    /* cfg,err:=ini.Load("./conf/config.ini")
     if err!=nil{
     	fmt.Printf("Load config file failed,err:%v",err)
	 }
	 addr:=cfg.Section("kafka").Key("address").String()
     topic=cfg.Section("kafka").Key("topic").String()
     path:=cfg.Section("taillog").Key("path").String()
     fmt.Println(addr,topic)
     fmt.Println(path)*/
	//cfg=new(conf.AppConf) //初始化结构体对象
	//调用MaoTo方法之前，一定要初始化结构体，否则出现空指针异常
	//当然，也可以在声明这个全局变量的时候初始化
	ini.MapTo(cfg,"./conf/config.ini")
	fmt.Println(cfg.KafkaConf.Address)
	//2.初始化kafka连接
	err:=kafka.Init([]string{cfg.KafkaConf.Address})

	if err != nil {
		fmt.Println("Init kafka failed, err:", err)
		return
	}
    fmt.Println("Init kafka successfully.....")
	//3.打开日志文件准备收集
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err!=nil{
		fmt.Println("Init taillog failed, err:", err)
		return
	}
	fmt.Println("Init taillog successfully.....")
    //4.业务逻辑
	run()
}

func run() {

	//1.读取日志

	for {
		select {
		case line :=<- taillog.ReadChan():
			//2.发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic,line.Text)
		default:
			time.Sleep(time.Second)
	   }
	}

}
