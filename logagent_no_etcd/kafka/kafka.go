package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/*func main() {
	nsqconsumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	// 根据topic取到所有的分区
	partitions, err := nsqconsumer.Partitions("web_log")
	if err!=nil{
		fmt.Printf("get partitions failed ,err:%v",err)
		return
	}
	fmt.Println(partitions)
	//遍历所有分区
	for partition:=range partitions{
		fmt.Printf("%v",partition)
		// 针对每个分区创建一个对应的分区消费者
		consumePartition, err := nsqconsumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
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
}*/
var (client sarama.SyncProducer)
//初始化生产者
func Init(addr []string)(err error){
   config:=sarama.NewConfig()
   config.Producer.Return.Successes=true
   config.Producer.Partitioner=sarama.NewRandomPartitioner
   config.Producer.RequiredAcks=sarama.WaitForAll

	// 连接kafka
	client,err=sarama.NewSyncProducer(addr, config)

	if err != nil {
		fmt.Println("nsqproducer closed, err:", err)
		return
	}
   return
}

func SendToKafka(topic,data string)  {
	// 构造一个消息
	msg:=&sarama.ProducerMessage{}
	msg.Topic=topic
	msg.Value=sarama.StringEncoder(data)

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err!=nil{
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}