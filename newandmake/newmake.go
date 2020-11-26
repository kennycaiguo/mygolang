package main

import "fmt"

func main() {
	/*var a *int//不要这样子定义指针变量
	*a=100
	fmt.Println(*a)*/  //这里会出错，因为没有给变量分配内存

	var a=new(int) //用new生成指针，会给它分配内存空间
	*a =100
	fmt.Println(*a) //100
}
