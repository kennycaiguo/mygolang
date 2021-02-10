package taillog

import "kenny.com/logagent/etcd"

var (
	tskMgr *tailLogMgr
)
type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	//tskMap map[string]*TailTask
}

func Init(logEntryConf []*etcd.LogEntry)  {
   tskMgr=&tailLogMgr{
   	logEntry: logEntryConf,
   }
	for _,logentry:=range logEntryConf{
		NewTailTask(logentry.Path,logentry.Topic)
	}
}