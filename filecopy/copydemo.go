package main

import (
	"io"
	"os"
)

func FileCopy(srcname ,destname string){
	src,_:=os.Open(srcname)
	defer src.Close()
	dest,_:=os.OpenFile(destname,os.O_CREATE|os.O_WRONLY,0666)
	defer dest.Close()
	io.Copy(dest,src)
}
func main() {
     file:="install.txt"
     FileCopy(file,"new.txt")
}
