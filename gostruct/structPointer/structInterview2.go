package main

import "fmt"

type student2 struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student2)
	stus := []student2{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	 for i:=0;i< len(stus);i++{
	 	m[stus[i].name]=&stus[i]
	 }

	//fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		//小王子 => 小王子
		//娜扎 => 娜扎
		//大王八 => 大王八

	}
}