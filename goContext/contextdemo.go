package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var exit bool

func job()  {
	for{
		fmt.Println("i love girls,i want get one....")
		time.Sleep(time.Second);
		if exit{
			break
		}
	}
	wg.Done()
}
func jobchan(exitChan chan struct{}){
	LOOP:
	for{
		fmt.Println("i love girls,i want get one....")
		time.Sleep(time.Second);
		select {
		case <- exitChan:
			break LOOP
		default:

		}
	}
}
func jobchan2(exitChan chan bool){
LOOP:
	for{
		fmt.Println("i love girls,i want get one....")
		//time.Sleep(time.Second);
		select {
		case <- exitChan:
			break LOOP
		default:

		}
	}
}

func jobContext(ctx context.Context)  {
	go contexwk(ctx)
LOOP:
	for{
		fmt.Println("i love girls,i want get one....")
		time.Sleep(time.Second);
		select {
		case <- ctx.Done():
			break LOOP
		default:

		}
	}
	wg.Done()
}
func contexwk(ctx context.Context)  {
	LOOP:
	for{
		fmt.Println("i love women,i am going to get one....")
		time.Sleep(time.Second);
		select {
		case <- ctx.Done():
			break LOOP
		default:

		}
	}
	wg.Done()
 }

func main01() {
	//var exitChan = make(chan bool)
	var exitChan2 = make(chan struct{})
   wg.Add(1)
  /* go jobchan2(exitChan)
   time.Sleep(time.Second*5)
	exitChan <- true*/
	go jobchan(exitChan2)
	time.Sleep(time.Second*5)
	exitChan2 <- struct{}{}
   wg.Wait()
}

func main() {
   wg.Add(1)
   ctx,cancel:=context.WithCancel(context.Background())
	go jobContext(ctx)
   time.Sleep(time.Second*10)
   cancel()
   wg.Wait()

}