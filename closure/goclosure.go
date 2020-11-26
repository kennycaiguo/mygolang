package main

import (
	"fmt"
	"strings"
)

//定义闭包
func adder() func(int) int { //就是函数的返回值也是一个函数
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

//闭包进阶示例2：比较难理解，相当于有内外两层函数，把内层函数作为外层函数的返回值，先从内层函数开始执行，内层函数没有的参数会往外层找
//最后把结果返回
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func calc(base int) (func(int)int,func(int)int){
	add:= func(x int) int{
		base+=x
		return base
	}
	sub:= func(x int) int {
		base-=x
		return base
	}
	return add,sub
}

func main() {
	//匿名函数就是没有名字的函数，在一个函数的内部只能定义匿名函数
	/*add:= func(a,b int) int {
		return a+b
    }
    fmt.Println(add(10,20))*/

	//自执行函数：匿名函数定义完加()直接执行
	/*func(x, y int) {
		fmt.Println(x + y)
	}(100, 20)*/

	//闭包参数
	/*f:=adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60
	fmt.Println(f(40)) //100*/

	/*jpgFunc := makeSuffixFunc(".jpg")//赋值时的格式是外层函数的格式
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg //执行时的参数是内层函数的格式
	fmt.Println(txtFunc("test")) //test.txt*/
	//以下这个叫做立即执行函数，它只能使用一次
	/*func(x,y int){
		fmt.Printf("x=%d,y=%d\n",x,y) //func。。。{}为定义函数
	}(10,12) //后面的（int，int）为调用函数*/
   f1,f2:=calc(10)
   fmt.Println(f1(1),f2(2))//11,9
   fmt.Println(f1(3),f2(4))//12,8
   fmt.Println(f1(5),f2(6))//13,7
}
