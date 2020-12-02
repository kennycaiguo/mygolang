package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int){
 defer wg.Done()
 fmt.Printf("gorountine %d\n",i)

}
 //sync 包的使用
func main() {//main函数本身就是一个主goroutinein


	for i:=0;i<10;i++ {
		wg.Add(1)
		go hello(i)

	}
     //开启一个goroutine去执行say函数
    wg.Wait()
     println("main")
     //go say()
}