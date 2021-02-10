package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"kenny.com/logagent/kafka"
)

//var (
//	tailObj *tail.Tail
//	LogChan chan string
//
//)
//TailTask是一个日志收集的任务
type TailTask struct {
	path string
	topic string
	instance *tail.Tail
}

func NewTailTask(path,topic string) (tailObj *TailTask) {
  tailObj=&TailTask{
  	 path: path,
  	 topic:topic,
  }
  tailObj.init() //根据路径打开对应的日志
  return
}

func (t *TailTask)init()  {
	config:=tail.Config{
		ReOpen: true,//重新打开
		Follow: true,//是否跟随
		Location: &tail.SeekInfo{Offset: 0,Whence: 2},//从什么位置开始读
		MustExist: false,//文件不存在不报错
		Poll: true,//
	}

	var err error
	t.instance,err=tail.TailFile(t.path,config)
	if err!=nil{
		fmt.Println("tail file err=", err)
	}
	go t.run() //收集日志发送到kafka
}

func (t *TailTask) run()  {
	for{
		select {
		case line:=<-t.instance.Lines: //从tailObj的通道在逐行读取数据
			//4.2 发往kafka
			//kafka.SendToKafka(t.topic,line.Text)
			//优化
			//先把日志放到一个通道中
			kafka.SendToChan(t.topic,line.Text)
			//kafka包中有一个单独的goroutine去取日志数据发到kafka

		default:

		}
	}
}
