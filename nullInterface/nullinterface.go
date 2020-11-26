package main

import "fmt"

func main() {
	var x interface{} //这是一个空接口，它可以接收任何类型的变量
	s:="Hello sexy girl"
	x=s
	fmt.Printf("%T\n" ,x)
	a:=[]int{1,2,3}
	for i,v:=range a{ //可以遍历数组
		fmt.Printf("%d=>%v\n",i,v)
	}
	x=a
	//fmt.Printf("%T\n",x)
/*	for i,v:=range x{ //错误，不能遍历接口，就算它接收了数组
		fmt.Println("%d=>%v",i,v)
	}*/
}
