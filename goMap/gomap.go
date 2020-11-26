package main

import "fmt"

func main() {
	m:=map[string]int{"age":20,"height":175} //创建map对象并且赋值
	println(m) //0xc00003be78
    fmt.Println(m) //map[age:20 height:175] 所以要用fmt.Println
    fmt.Println(m["age"]) //20
	for s := range m {
		fmt.Println(s)    //age "\n" height
		fmt.Println(m[s]) //20 "\n" 175
	}

}
