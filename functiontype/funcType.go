package main

import "fmt"

type calculation func(int,int) int //声明一个函数类型，该类接收两个整数参数，有一个整型返回值，凡是符合这个条件的函数都可以给它赋值

func add(x, y int) int{
	return x+y
}
func sub(x, y int) int{
	return x-y
}
func div(x, y int) int{
	return x/y
}
func mul(x, y int) int{
	return x*y
}

func test(cal calculation) int{ //函数的参数也可以是函数类型
	cal =add
   return	cal(12,23)
}
func main() {
  var c calculation //声明函数类型变量
  /*c=add   //将指定的函数赋值给函数变量
  fmt.Println(c(12,22))*/ //就可以利用变量名加参数来调用这个函数，有点像c#中的委托
  /*c=mul
  fmt.Println(c(12,22))
  fmt.Printf("type of c:%T\n", c)*/
  fmt.Println(test(c))
}
