package main


import (
	"net/http"
	"io"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "你好，client。。")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `
        <!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width = device-width, initial-scale = 1.0">
        <meta http-equiv="X-UA-Compatible" content = "ie=edge">
        <title> 首页</title>
        <style>
            body{
                background-image: url(hot.jpg)
            }
        </style>
    </head>
    
    
    <body>
        <h1>welcome to index</h1>
    </body>
</html>

`)
	})

	mux.HandleFunc("/baidu", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://www.baidu.com", http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/sayhi", sayhello)
	http.ListenAndServe(":8080", mux)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}