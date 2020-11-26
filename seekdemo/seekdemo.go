package main

import (
	"io/ioutil"
	"os"
)

func insertcontent(n int,content string,path string)  {
	f,_:=os.OpenFile(path,os.O_RDWR,0666)
	str, _ := ioutil.ReadFile(path)
	start:=str[:n+1]
	remain:=str[n:]
	strStart:=string(start)
	strRemain:=string(remain)
	strAll:=strStart+content+strRemain
	f.WriteString(strAll)
    f.Close()
}
func main() {
	/*f,_:=os.OpenFile("hello.txt",os.O_RDWR,0666)
	f.Seek(36,0)
	/*var tem [20]byte
	n,_:=f.Read(tem[:])
	fmt.Println(string(tem[:n]))
	f.WriteString("hello,sexy")*/
   insertcontent(61,"hello,vaginas!!!","hello.txt")
}
