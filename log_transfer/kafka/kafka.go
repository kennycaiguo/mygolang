package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"kenny.com/log_transfer/es"
)
type LogData struct {
	Data string`json:"data"`
}
func Init(addrs []string,topic string)(error){
	//创建分区消费者
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			    //直接发给ES

			   ld:=LogData{
			   	Data: string(msg.Value),
			   }
				/*if err2 != nil {
					fmt.Printf("json Unmarshal failed,err:%v",err2)
					continue
				}*/
				es.SendToES(topic,ld)

			}
		}(pc)

	}
   return err
}

