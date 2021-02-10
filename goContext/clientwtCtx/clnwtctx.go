package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

/**
  在client中使用context
 */
//创建一个结构体用于接收服务器返回结果
type RespData struct {
	resp *http.Response
	err error
}
//创建一个使用context的请求函数
func sendReq(ctx context.Context){
	//创建transport对象
	transport:=http.Transport{
		// 请求频繁可定义全局的client对象并启用长链接
		// 请求不频繁使用短链接
		DisableKeepAlives: true,
	}
	//创建客户端
	client:=http.Client{
		Transport: &transport,
	}
	//创建返回数据通道
	respChan:=make(chan *RespData,1)
	req,err:=http.NewRequest("get","http://127.0.0.1:8000",nil)
	if err!=nil{
		fmt.Printf("new requestg failed, err:%v\n", err)
		return
	}
	//原请求对象利用context创建新的请求对象
	req=req.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	//创建匿名goroutine
	go func() {
		//发送请求
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		//将返回的数据保存到RespData中
		rd:=&RespData{
			resp:resp,
			err:err,
		}
		//将其放入通道
		respChan<-rd
		wg.Done()
	}()
	//利用select来处理
	select {
	case <- ctx.Done():

	case result:=<-respChan:
       fmt.Println("send request successfully....")
       if result.err!=nil{
		   fmt.Printf("call server api failed, err:%v\n", result.err)
		   return
	   }
	   //接收数据并且输出
	   data,_:=ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}
func main() {
   // 定义一个100毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel() // 调用cancel释放子goroutine资源
	sendReq(ctx)

}
