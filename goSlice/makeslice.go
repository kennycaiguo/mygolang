package main

import "fmt"

//演示make函数的使用
func main(){
	a:=make([]int,3,10)//创建一个长度为3，容量为10的切片
	s:=[]int{1,4,6}
	a=s //切片拷贝，其实就是内存地址的传递
	fmt.Println(a,s)
	s[1]=40 //切片是引用类型，切片其实不保存值，而是有底层数组来保存
	fmt.Println(a,s)



}
