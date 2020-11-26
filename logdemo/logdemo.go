package main

import (
	"log"
	"os"
)

func logTofile(path string)  {
	file,_:=os.OpenFile(path,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	//fmt.Fprintf(os.Stdout,"this is a log message\n")
	log.SetOutput(file)
	//log.SetFlags(log.Llongfile|log.Lmicroseconds|log.Ldate)
	log.SetFlags(log.Lshortfile|log.Lmicroseconds|log.Ldate)
	log.Println("this is a log message\n")
	log.SetPrefix("[Info]")
	log.Println("this is a log message\n")
}
func logToStdOut()  {
	logger:=log.New(os.Stdout,"[Info]",log.Lshortfile|log.Lmicroseconds|log.Ldate)
	logger.Println("this is a common log message")
}
func main() {
  //logTofile("log.log")
   logToStdOut()

}
