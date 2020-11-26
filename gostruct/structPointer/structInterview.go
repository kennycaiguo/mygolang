package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus { //这里将所有的key的值都设置位相同的地址
		m[stu.name] = &stu
	}
	//fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		/*小王子 => 大王八  //想一想为什么？
		娜扎 => 大王八
		大王八 => 大王八*/

	}
}