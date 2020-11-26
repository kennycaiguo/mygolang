package main

import (
	"fmt"
)

func main()  {
	var s []int //这样子创建的切片是空的，没有长度也没有容量
	a:=make([]int,5,10) //这样子创建的切片，有长度也有容量，要注意区分
	for i:=0;i<10;i++{
		a = append(a,i)//长度为5，就从第六的位置追加
		s = append(s,i)//长度为0，从索引0处追加

	}
	fmt.Println(a) //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
    fmt.Println(s) //[0 1 2 3 4 5 6 7 8 9]

}
