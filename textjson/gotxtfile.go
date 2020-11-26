package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//writetxtfile()
	//writeTextfile2()
	ioutilWrite("./test.txt")
    //readTextFile("./hello.txt")
    //readFileBufio("./hello.txt")
	//readWithIoUtil("./hello.txt")
}

func ioutilWrite(s string) {
	tem:=[]byte("this is a test")
	ioutil.WriteFile(s,tem,0666)

}

func readWithIoUtil(s string) {
	f,e:=os.OpenFile(s,os.O_RDONLY,0666)
	if e!=nil{
		fmt.Printf("open error:%s",e)
	}
	contents, err := ioutil.ReadAll(f)
	if err!=nil{
		fmt.Println("read error:%s",err)
	}
	defer f.Close()
	fmt.Printf("content is :%s",contents)

}

func readFileBufio(s string) {
	var tmp byte
	f,_:=os.OpenFile(s,os.O_RDONLY,0666)
	defer f.Close()
	for{
		reader:=bufio.NewReader(f)
		str, err:= reader.ReadString(tmp)

		fmt.Printf("content:%s\n",str)
		if err==io.EOF{
			fmt.Println("finished...")
			return
		}
	}


}

func readTextFile(path string) {
	f,_:=os.OpenFile(path,os.O_RDONLY,0666)
	defer f.Close()
	var content []byte
	var tmp = make([]byte,128)
	for{
		n,err:=f.Read(tmp)
		if err==io.EOF{
			fmt.Println("read finished")
			break
		}
		content = append(content,tmp[:n]...)
	}
    fmt.Printf("file content:%s",string(content))
}

func writetxtfile() {
	str := "what a nice day!!! how are u feeling"
	f, err := os.OpenFile("hello.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	n, err := fmt.Fprint(f, str)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("写入的size：%d\n", n)



	f.Close()
}
func writeTextfile2(){
	age:=20
	height:=175
	gender:="male"
	name:="Benny choi"
	words:="i believe i can fly ,if iam in a plane!!!他是个大sb"
	info:=fmt.Sprintf("age:%d,height:%dcm,gender:%s\n name:%s,words：%s",age,height,gender,name,words)
	f,_:=os.OpenFile("structPerson.txt",os.O_CREATE|os.O_WRONLY,0666)
	len,_:= fmt.Fprint(f,info)
	fmt.Printf("写入的size：%d",len)
	f.Close()
}

