package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"kenny.com/log_transfer/conf"
	"kenny.com/log_transfer/es"
	"kenny.com/log_transfer/kafka"
)

/**
log transfer，功能：
将日志数据kafka取出来，放入es里面

 */
func main() {
	  var cfg conf.LogTransferCfg
  //加载配置文件
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init config failed ,err:%v",err)
		return
	}
	fmt.Println(cfg)
	//1.初始化es
	err = es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Printf("init elasticsearch failed,err:%v",err)
		return
	}
	//2.初始化kafka，并且往es发送数据
	/*if !strings.HasPrefix(cfg.KafkaCfg.Address,"http://"){
		cfg.KafkaCfg.Address="http://"+cfg.KafkaCfg.Address
	}
	fmt.Println(cfg.KafkaCfg.Address)*/
   err=kafka.Init([]string{cfg.KafkaCfg.Address},cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka failed,err:%v",err)
		return
	}

	select {}
}
