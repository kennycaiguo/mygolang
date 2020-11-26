package main

import (
	"reflect"
	"fmt"
)

type person struct {
	Name   string
	Gender string
	Age    int
	Address string
	Phone string
    Hobby []string
	Email string
}

func Struct2Map(obj interface{}) map[string]interface{} {
	     t := reflect.TypeOf(obj)
	    v := reflect.ValueOf(obj)

	     var data = make(map[string]interface{})
	    for i := 0; i < t.NumField(); i++ {
		          data[t.Field(i).Name] = v.Field(i).Interface()
		     }
	     return data
}

func main() {
    var p person //创建结构体变量
    p.Name="Jackline"
    p.Gender="female"
    p.Age = 20
    p.Address="3 pawsey road kgn5"
    p.Phone="87716563"
    p.Email="Jackline@gmail.com"
    p.Hobby =[]string{"football","swimming","reading","cooking","sex"}
    mapperson:=Struct2Map(p)
    for k,v:=range mapperson{
		fmt.Printf("%v : %v\n",k,v)
	}

}
