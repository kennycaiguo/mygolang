package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
/**
WithValue返回父节点的副本，其中与key关联的值为val。

仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。

所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。
 */
type KeyCode string

var wg3 sync.WaitGroup

func ctxWtValue(ctx context.Context)  {
   //创建一个key
	key:=KeyCode("KeyCode")
	// 在子goroutine中获取keyCode
	keyCode,ok:=ctx.Value(key).(string)
	if !ok{
		fmt.Println("invalid key !!!")
	}
	LOOP:
		for{
			fmt.Printf("key code:%s\n",keyCode)
			time.Sleep(time.Millisecond*10)
			select {
			case <- ctx.Done():
				break LOOP
			default:

			}
		}
	fmt.Println("worker done!")
	wg3.Done()
}
func main() {
	// 设置一个50毫秒的超时
   ctx,cancel:=context.WithTimeout(context.Background(),time.Millisecond*50)
	// 在系统的入口中设置keycode传递给后续启动的goroutine实现日志数据聚合
	ctx=context.WithValue(ctx,KeyCode("KeyCode"),"12345678")
	wg3.Add(1)
	go ctxWtValue(ctx)
	time.Sleep(time.Second * 5)
	//通知子goroutine结束
	cancel()
	wg3.Wait()
	fmt.Println("over")
}