package main

import "fmt"

func main() {
	   ch:=make(chan int,1)
	   for i:=0;i<10;i++{
		   select {
		   case x:=<-ch:
			   fmt.Printf("x=%d\n",x) //0 2 4 6 8
		   case ch<-i:

		   }
	   }
	}
/*
以上代码运行原理：当通道没有值，会走发送哪个case，将零发送到通道，送完后，i、变成1，
因为这里通道的缓冲区只能存放一个数字，所以发送完后，只能取值，所以将0取出，显示在屏幕上，
此时i的值已经是2了，然后又只能发送，就将2发送到通道，。。。一直重复以上过程，所以
就得到0，2，4，6，8的结果
*/