package main

import (
	"fmt"
	"time"
)

func say()  {
	fmt.Println("hello ,goroutine")
}

func main() {//main函数本身就是一个主goroutinein
	for i:=0;i<10;i++ {
		go func(i int) {
           fmt.Println(i)
		}(i)
	}
     //开启一个goroutine去执行say函数
    time.Sleep(1000000)//如果不写这一句，就等不到新创建的goroutine执行，因为main函数执行完成，整个程序就会退出
     println("main")
     //go say()
}