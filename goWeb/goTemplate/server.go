package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func tmpl(w http.ResponseWriter, r *http.Request)  {
	tpl,err:=template.ParseFiles("test.html")
	if err!=nil{
		fmt.Println(err.Error())
	}

	tpl.Execute(w,"我是后台传过来的内容")
}
func main()  {
	server :=http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/test",tmpl)
	server.ListenAndServe()
}
