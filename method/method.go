package main

import "fmt"

type myint int32 //1.利用type关键字从基本数据类型中定义一个自定义类型
//go语言可以给任意自定义数据类型添加方法，这里我给类添加
func (i myint) power(x myint) (res myint){ //给自定义类型添加方法

	return x*x
}


func (i myint)mySqrt1(x myint) myint {
	res:= x
	//牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}


func main() {
	/*var a myint=3
	fmt.Printf("%d 的平方是%d",a,a.power(a))*/
	var b myint=625
	fmt.Printf("%d的平方根是%d",b,b.mySqrt1(b))
}
