package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)
type logData struct {
	topic string
	data string
}
var (
	client sarama.SyncProducer
	logDataChan chan *logData
	)
//初始化生产者
func Init(addr []string,maxSize int)(err error){
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
	//初始化logDataChan
	logDataChan=make(chan *logData,maxSize)
	//开启goroutine从后台获取日志信息发送到kafka
	go sendToKafka()
   return
}

func SendToChan(topic,data string){
	ldmsg:=&logData{
		topic: topic,
		data: data,
	}
	logDataChan<-ldmsg
}
//真正往kafka发送日志的函数
func sendToKafka()  {
	for  {
		select {
		case ld:=<- logDataChan:
			// 构造一个消息
			msg:=&sarama.ProducerMessage{}
			msg.Topic=ld.topic
			msg.Value=sarama.StringEncoder(ld.data)

			//发送消息
			pid, offset, err := client.SendMessage(msg)
			if err!=nil{
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond*50)
		}

	}



	//

}

//将日志数据发送到通道
