package main

import "fmt"

func main() {
	//deferdemo1()
	deferdemo2()
}

func deferdemo2() {
	fmt.Println("start")
	defer fmt.Println("hahahaha....") //defer 的作用：将这个语句推迟到即将返回的时候执行
	fmt.Println("end")
}

func deferdemo1() {
	fmt.Println("start")
	defer fmt.Println(1) //注意：defer的机制和栈有点相似，都是最后的语句最先执行
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
