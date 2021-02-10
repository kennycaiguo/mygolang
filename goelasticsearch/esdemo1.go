package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Printf("Init client failed,error:%v",err)
		panic(err)
	}

	fmt.Println("connect to es success")
	p:=Person{
		Name:"Jack",
		Age: 28,
		Married: false,
	}
	put1, err := client.Index().Index("user").BodyJson(p).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n",put1.Id,put1.Index,put1.Type)
}
