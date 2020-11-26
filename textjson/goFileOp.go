package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {

	filePath := "../txt/myfile.txt" //txt目录必须存在
	//创建文件
	//writeFile(filePath)
	writeFile2(filePath)
	//readFile(filePath)

}

func writeFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666) //目录必须存在否则报错
	if err != nil {
		fmt.Println("打开文件失败！！！")
	}
	//关闭文件句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i <= 3; i++ {
		writer.WriteString("hello ,i am created by golang!\n")
	}
	writer.Flush()
}

func writeFile2(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666) //目录必须存在否则报错
	if err != nil {
		fmt.Println("打开文件失败！！！")
	}
	//关闭文件句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString("hello ,this file iscreated by golang!\n")
	writer.WriteString("golang is a young young programming language\n")
	writer.WriteString("i am learning golang!\n")
	writer.WriteString("golang has a very bright future!!!\n")

	writer.Flush()
}

func readFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err!=nil {
		fmt.Println("读取文件失败")
	}
	defer file.Close()
	reader:=bufio.NewReader(file)

	for {
		str,err := reader.ReadString('\n')
        if err == io.EOF{
        	break
		}
        fmt.Println(str)
	}
}