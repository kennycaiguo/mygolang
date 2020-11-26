package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main(){
	//writeGobFile()
	readGobFile()
}

func readGobFile() {
	hfile ,err:= os.Open("output.gob")
	if err!=nil{
		fmt.Println("打开文件失败")
		return
	}
	defer hfile.Close()
	info:= ""
    decoder:= gob.NewDecoder(hfile)
    err = decoder.Decode(&info)
	if err!=nil{
		fmt.Println("解码失败:%v",err)
	}else{
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

func writeGobFile() {
	info := "http://c.biancheng.net/golang/"
	hfile ,err:= os.Create("output.gob")
	if err!=nil{
		fmt.Println("创建文件失败")
		return
	}
	defer hfile.Close()

	encoder:=gob.NewEncoder(hfile)

	err = encoder.Encode(info)
	if err!=nil{
		fmt.Println("编码失败")
	}else{
		fmt.Println("编码成功")
	}
}
