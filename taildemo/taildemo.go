package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tailf的用法示例

func main() {
   fileName:="./mylog.log"
   //首先初始化配置结构体config
   config:=tail.Config{
   	   ReOpen: true,
   	   Follow: true,
   	   Location: &tail.SeekInfo{Offset: 0,Whence: 2,},
   	   MustExist: false,
   	   Poll: true,
   }
   //用TailFile函数，并传入文件路径和config，返回有个tail的结构体，tail结构体的Lines字段封装了拿到的信息
	tailFile, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line
		ok bool
	)
	//遍历tail.Lnes字段，取出信息（注意这里要循环的取，因为tail可以实现实时监控）
	for{
		line,ok = <- tailFile.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tailFile.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
