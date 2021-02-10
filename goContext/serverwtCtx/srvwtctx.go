package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)
/**
 服务器没有使用到contex
 */
//编写用户请求的处理函数，对应路径是“/”
func indexHandler(w http.ResponseWriter, r *http.Request)  {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10) // 耗时10秒的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}
func main() {
   //添加对“/”处理函数的映射
	http.HandleFunc("/",indexHandler)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err!=nil{
		panic(err)
	}
}
