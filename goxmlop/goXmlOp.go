package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main(){
  //writeXml()
  readXml()
}

func writeXml() {
	//实例化对象
	info := Website{"C语言中文网", "http://c.biancheng.net/golang/", []string{"Go语言入门教程", "Golang入门教程"}}

	hFile,err := os.Create("info.xml")
	if err!=nil{
		fmt.Println("创建文件失败",err.Error())
		return
	}
	defer hFile.Close()
    encoder:= xml.NewEncoder(hFile)
    err = encoder.Encode(info)
    if err!=nil{
		fmt.Println("编码失败")
	}else{
		fmt.Println("编码成功")
	}
}

func readXml(){
	hfile,err :=os.Open("info.xml")
	if err!=nil{
		fmt.Println("打开文件失败")
		return
	}

	defer hfile.Close()
	info :=Website {}
	decoder:=xml.NewDecoder(hfile)

	err = decoder.Decode(&info)

	if err!=nil{
		fmt.Println("解码失败:%v",err)
	}else{
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}