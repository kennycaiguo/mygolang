package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	// 根据topic取到所有的分区
	partitions, err := consumer.Partitions("web_log")
	if err!=nil{
		fmt.Printf("get partitions failed ,err:%v",err)
		return
	}
	fmt.Println(partitions)
	//遍历所有分区
	for partition:=range partitions{
		fmt.Printf("%v",partition)
		// 针对每个分区创建一个对应的分区消费者
		consumePartition, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err!=nil{
			fmt.Printf("failed to start nsqconsumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer consumePartition.AsyncClose()
		// 异步从每个分区消费信息
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg:=range consumePartition.Messages(){
				   fmt.Printf("Partition:%v,Offset:%v,key:%v,value:%v", msg.Partition,msg.Offset,msg.Key,msg.Value)
				}
		}(consumePartition)

	}
}