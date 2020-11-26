package main

import "fmt"

func main(){
	/* s:=[]int{1,2,3}
     fmt.Println(s)
	//ints := append(s, 10)//append函数其实并没有子标切片s，它生成一个扩展后的新切片，当然也可以用s来接收
	s= append(s, 10)//append函数其实并没有子标切片s，它生成一个扩展后的新切片，当然也可以用s来接收
	fmt.Println(s)*/

	arr:=[4]int{11,22,33,44}
	a:=arr[:4]
	fmt.Println(a,cap(a))
	a = append(a, 4)
	fmt.Println(a,cap(a))
	a = append(a, 44) //如果扩容时，底层数组长度不够，而且容量小于1024字节append会把底层数组的长度增加一倍
	fmt.Println(a,cap(a))
	a = append(a, 55,66,77)//可以一次追加几个值
	fmt.Println(a)
	s:=[]int{10,20,30}
	a = append(a,s...) //注意，可以使用...将切片拆分为单给元素
	fmt.Println(a)
	//append的特殊用法删除“breast”元素，首先以要删除的元素为条件拆分处两个切片，然后再吧第二个切片拆开后append到第一个切片中
	b:=[]string{"pussy","vagina","breast","labias"}
    c:=make([]string,5,5)
    c=b[:]
    fmt.Println(b,c)
    c=append(c[:2],c[3:]...)
	fmt.Println(b,c)  //[pussy vagina labias labias] [pussy vagina labias]

}
