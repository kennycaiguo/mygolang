package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Website struct {
	Name string `xml:"name,ptr"`
	Url string
	Course []string

}
func main(){
   //writeJson();
   readJson()
}

func writeJson()  {
	//json数据
	info := []Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}

	//创建文件
	filehd,err := os.Create("info.json")
	if(err!=nil){
		fmt.Println("创建文件失败",err.Error())
		return
	}
	defer  filehd.Close()
	//创建Json编码器
	encoder:=json.NewEncoder(filehd)

	err = encoder.Encode(info)

	if err!=nil{
		fmt.Println("编码失败")
	}else{
		fmt.Println("编码成功")
	}

}

func readJson() {
	filehd,err := os.Open("info.json")
	if(err!=nil){
		fmt.Println("打开文件失败",err.Error())
		return
	}
	defer  filehd.Close()
	var info []Website
	decoder :=json.NewDecoder(filehd)

	err = decoder.Decode(&info)
	if err!=nil{
		fmt.Println("解码失败")
	}else{
		fmt.Println("解码成功")
		fmt.Println(info)
	}

}