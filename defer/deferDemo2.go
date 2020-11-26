package main

import "fmt"

func f1() int { //这是没有名称的返回值
x := 5
defer func() {
x++
}()
return x //5 这个就是给返回值赋值，以后执行defer语句，对返回值已经没有影响了
}

func f2() (x int) { //这是命名返回值，返回值用下接收，所以defer语句修改x，也就修改了返回值
defer func() {
x++
}()
return 5 //6
}

func f3() (y int) {
x := 5
defer func() {
x++
}()
return x //5
}
func f4() (x int) {
defer func(x int) {
x++
}(x)  //函数传参数修改的副本
return 5 //5
}
func main() {
fmt.Println(f1())
fmt.Println(f2())
fmt.Println(f3())
fmt.Println(f4())
}

