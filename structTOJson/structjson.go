package main

import (
	"../gostruct/structPointer/structPerson"
	"encoding/json"
	"fmt"
)
type person structPerson.Person
func main() {
	p1:= person{
		Name:"jack",
		Age:18,
		Gender:"male",
		City:"niewyork",
	}
	fmt.Println(p1);
	//序列化
	bytesjson, err := json.Marshal(p1)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(string(bytesjson)) //需要先将序列化的结果转为字符串，否则没有可读性
    //反序列化
	p2 := person{}
	json.Unmarshal(bytesjson,&p2)//第一个参数是字节类型，不能用字符串，如果是字符串必须转为字节类型
	fmt.Println(p2);

}
