package taillog

import (
	"fmt"
	"kenny.com/logagent_watch/etcd"
	"time"
)

var (
	tskMgr *tailLogMgr
)
//TailTask管理者
type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	tskMap map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry)  {
   tskMgr=&tailLogMgr{
   	logEntry: logEntryConf,
   	tskMap: make(map[string]*TailTask,16),
   	newConfChan: make(chan []*etcd.LogEntry),//这里不需要缓冲区
   }
	for _,logentry:=range logEntryConf{
		//初始化的时候起了多少个tailtask，都要记录下来，为了后续判断方便
		tailObj:=NewTailTask(logentry.Path,logentry.Topic)
		mk:=fmt.Sprintf("%s_%s",logentry.Path,logentry.Topic)
		tskMgr.tskMap[mk]=tailObj
	}
	go tskMgr.run()
}

//这个方法监听newConfChan,有变化就做处理
//包含
func (t *tailLogMgr) run()  {
	for  {
		select {
		case newConf:=<-t.newConfChan:
			for _,conf:=range newConf{
				mk:=fmt.Sprintf("%s_%s",conf.Path,conf.Topic)
               _,ok:=t.tskMap[mk]
               if ok{
               	//原来有了，不需要操作
               	continue
			   } else{
				   //新增，
			   	tailObj:=NewTailTask(conf.Path,conf.Topic)
				   //mk:=fmt.Sprintf("%s_%s",conf.Path,conf.Topic)
				   t.tskMap[mk]=tailObj
			   }

			}

			//找出原来t.logEntry里面有，newConf里面没有的，删除掉
            for _,c1:=range t.logEntry{
            	isDelete:=true
            	for _,c2:=range newConf{
                   if c2.Path==c1.Path && c2.Topic==c1.Topic{
                   	isDelete=false
                     continue
				   }
				}
				if isDelete{
					//把c1对应的tailObj停掉
					mk:=fmt.Sprintf("%s_%s",c1.Path,c1.Topic)
					t.tskMap[mk].cancelfunc()
				}
			}
			//删除配置
			fmt.Println("发现新配置....",newConf)
		default:
			time.Sleep(time.Second)
	}
  }
}

//定义一个公共函数返回一个newConfChan对象供外部调用
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}