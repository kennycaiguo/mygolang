package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"

	"os"
	"strings"
)

var producer *nsq.Producer

var err error

func initProducer(str string) (error) {
	config:=nsq.NewConfig()
	producer, err= nsq.NewProducer(str, config)
	if err!=nil{
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func main() {
	nsqAddress := "127.0.0.1:4150"
    err:=initProducer(nsqAddress)
	  if err!=nil{
		fmt.Printf("init failed err:%v",err)
	  }
	  //接收键盘输入
	  reader:=bufio.NewReader(os.Stdin)
	for  {
		data, err := reader.ReadString('\n')
		if err!=nil{
			fmt.Printf("read string from stdin failed, err:%v\n", err)
			continue
		}
		data=strings.TrimSpace(data)
		if strings.ToUpper(data) == "Q" { // 输入Q退出
			break
		}
		//相topic_demo发布数据
		err = producer.Publish("topic_demo", []byte(data))
		if err!=nil{
			fmt.Println(err.Error())
			continue
		}
	}
}
