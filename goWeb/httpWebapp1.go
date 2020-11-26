package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter,r *http.Request)  {
  fmt.Println("----------------------------------------")
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("Path:",r.URL.Path)
  fmt.Println("schema:",r.URL.Scheme)
  fmt.Println(r.Form["url_ long"])
  for k,v :=range r.Form{
  	fmt.Println("key",k)
  	fmt.Println("value",strings.Join(v,""))
  	//输出到客户端

  }
	fmt.Fprintf(w,"Hello ,welcome to go web page!!!")
}

func main() {
	//设置路由
	http.HandleFunc("/",hello)
	fmt.Println("server is ready: http://127.0.0.1:8880/")
	//监听端口
	err := http.ListenAndServe(":8880",nil)
	if err!=nil{
		log.Fatal("ListenAndServe: ", err)
	}

}