package main

import (
	"log"
	"net/http"
)

func main()  {
	err:=http.ListenAndServe(":2000",http.FileServer(http.Dir("fileserver/")))
	if err != nil {
		log.Fatal(err)
	}
}
