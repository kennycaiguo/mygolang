package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var consumer *nsq.Consumer
var err error

type  MyHandler struct {
     Title string
}

func (m *MyHandler) HandleMessage(msg *nsq.Message) error {
	addr := msg.NSQDAddress
	message := string(msg.Body)
	fmt.Println(m.Title,addr, message)
	return nil
}

func initConsumer(topic string, channel string, address string)(error){
	config:=nsq.NewConfig()
	//设置监测间隔
	config.LookupdPollInterval = 15 * time.Second
	consumer, err = nsq.NewConsumer(topic, channel, config)
	if err!=nil{
		fmt.Println(err.Error())
		return err
	}
	//var handler nsq.Handler
	handler:=&MyHandler{
		"good receiver",
	}
	//handler.HandleMessage(messa)
	consumer.AddHandler(handler)
	err= consumer.ConnectToNSQLookupd(address)
	if err!=nil{
		fmt.Printf(err.Error())
	}
	return nil
}
func main() {
	err=initConsumer("topic_demo","first","127.0.0.1:4161")
	if err!=nil{
		fmt.Printf("init nsqconsumer failed, err:%v\n", err)
		return
	}
	ch:=make(chan os.Signal)//定义一个信号通道
	signal.Notify(ch, syscall.SIGINT)//转发键盘中断信号到ch
	<-ch

}
