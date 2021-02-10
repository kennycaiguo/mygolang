package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	//创建一个配置对象
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	//config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Partitioner=sarama.NewManualPartitioner
	config.Producer.Return.Successes=true
	// 构造一个消息
	msg:=&sarama.ProducerMessage{}
	msg.Topic="web_log"
	msg.Value=sarama.StringEncoder("This is a web log test msg...")
	// 连接kafka
	client,err:=sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)

	if err != nil {
		fmt.Println("nsqproducer closed, err:", err)
		return
	}
	defer client.Close()
	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err!=nil{
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
