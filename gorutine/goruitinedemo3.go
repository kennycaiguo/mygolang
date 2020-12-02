package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

 func randNum(){
 	rand.Seed(time.Now().UnixNano())
 	for i:=0;i<10;i++{
 		num1:=rand.Int()
 		num2:=rand.Intn(10)
 		fmt.Println(num1,num2)
	}
 }
 func speak(i int){
	// defer wg1.Done() //在这里一定要加defer
 	time.Sleep(time.Millisecond*time.Duration(rand.Intn(300)))
 	fmt.Println(i)
 	 wg1.Done()
 }
 //sync 包的使用
func main() {//main函数本身就是一个主goroutinein
	for i:=0;i<5;i++{
		wg1.Add(1)
		go speak(i)
  }
  wg1.Wait()
	//randNum()
}