package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg2 sync.WaitGroup
func gen(ctx context.Context) <-chan int {
	dst:=make(chan int)
	n:=1
	go func() {
		select {
		case <-ctx.Done():
			return
		case dst<-n:
			n++
		}
	}()
	return dst
}

func contextWithDealindemo(){
	d:=time.Now().Add(time.Millisecond*2000)//time.Millisecond*50 =>err,time.Millisecond*2000=>overslept
	ctxd,cancel:=context.WithDeadline(context.Background(),d)
	defer cancel()
	select {
	case <-time.After(time.Second):
		fmt.Printf("overslept...")
	case <-ctxd.Done():
		fmt.Println(ctxd.Err())
	}

}
func ctxwtTimeout(ctx context.Context)  {
	LOOP:
	for {
		fmt.Printf("logging in...\n")
		time.Sleep(time.Second)
		select {
		  case <-ctx.Done():
		  	break LOOP
		default:

		}
	}
	println("done")
	wg2.Done()
}
func main() {

  /* ctx,cancel:=context.WithCancel(context.Background())
   defer cancel()
  for n:=range gen(ctx){
  	fmt.Printf("n=%d\n",n)
  	if n==5{
  		break
	}
  }*/
//contextWithDealindemo()
//withTimeout
	wg2.Add(1)
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	ctxwtTimeout(ctx)
	wg2.Wait()
}