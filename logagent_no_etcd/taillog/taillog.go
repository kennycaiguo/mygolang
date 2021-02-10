package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

func Init(filename string)  (err error){
     config:=tail.Config{
     	ReOpen: true,//重新打开
     	Follow: true,//是否跟随
     	Location: &tail.SeekInfo{Offset: 0,Whence: 2},//从什么位置开始读
     	MustExist: false,//文件不存在不报错
     	Poll: true,//
	 }
	 tailObj,err=tail.TailFile(filename,config)
	 if err!=nil{
		 fmt.Println("tail file err=", err)
		 return
	 }
	return nil
}

func ReadChan() <-chan *tail.Line{
    return tailObj.Lines
}

func ReadLog()  {
	var (
		line *tail.Line
		ok bool
	)
	for{
		line,ok = <- tailObj.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tailObj.Filename)
			time.Sleep(time.Second)
			continue
		}
		 LogChan<-line.Text
	}
}