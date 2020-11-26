package main

import (
	"fmt"
	"./structPerson"
)


func newperson(name string,age int,gender,city string) *structPerson.Person{ //go语言中的结构体是没有构造函数的，我们来模拟一个
	return &structPerson.Person{
		name,age,gender,city,
	}
}
func main() {
	p:=newperson("Mike",18,"male","new york")
	fmt.Printf("%#v",p)
}
