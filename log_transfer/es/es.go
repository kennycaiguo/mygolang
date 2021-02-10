package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)
var (

	client *elastic.Client
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}
//初始化ES客户端，准备接受kafka发来的数据
func Init(addr string)(err error)  {
	if !strings.HasPrefix(addr,"http://"){
		addr="http://"+addr
	}
	client, err= elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		fmt.Printf("Init client failed,error:%v",err)
		return err
	}

	fmt.Println("connect to es success,addr:%v",addr)
	return
}

//发送到ES
func SendToES(index string,data interface{})(err error)  {
	//p:=Person{
	//	Name:"Jack",
	//	Age: 28,
	//	Married: false,
	//}
	put1, err := client.Index().Index(index).BodyJson(data).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n",put1.Id,put1.Index,put1.Type)
	return
}