package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)
var (

	client *elastic.Client
	ch chan *LogData

)
type LogData struct {
	Topic string`json:"topic"`
	Data string`json:"data"`
}
type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}
//初始化ES客户端，准备接受kafka发来的数据
func Init(addr string,chanSize,nums int)(err error)  {
	if !strings.HasPrefix(addr,"http://"){
		addr="http://"+addr
	}
	client, err= elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		fmt.Printf("Init client failed,error:%v",err)
		return err
	}

	fmt.Println("connect to es success,addr:%v",addr)
	ch=make(chan *LogData,chanSize)
	for i := 0; i <nums ; i++ {
		go SendToES()
	}

	return
}

//发送到ES

func SendToESChan(msg *LogData)  {
	 //
	ch <- msg

	//return
}
//func SendToES(index string,data interface{})(err error)  {
func SendToES() (err error) {
	 //
	for  {
		select {
		   case msg:=<-ch:
			   put1, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			   if err != nil {
				   fmt.Println(err)
				   continue
			   }
			   fmt.Printf("Indexed user %s to index %s, type %s\n",put1.Id,put1.Index,put1.Type)
		   default:
			time.Sleep(time.Second)
		}
	}

}