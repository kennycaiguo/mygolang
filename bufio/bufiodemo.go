package main
//go语言中出了可以使用fmt.Scanln方法获取键盘输入，还可以使用bufio.NewReader(os.Stdin)创建的reader对象的readLine方法
import (
	"bufio"
	"fmt"
	"os"
)

func bufioInput()  {
	//var s string
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容：\n")
	/*line, _, _ := reader.ReadLine()
	s = string(line)*/
	str, _ := reader.ReadString('\n')
	fmt.Printf("你输入的是：%s",str)
}
func writeFileWithbufioInput()  {
	path:="test.txt"
	reader:=bufio.NewReader(os.Stdin)
	file,_:=os.OpenFile(path,os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
	for{
		fmt.Println("请输入内容:")
		str, _:= reader.ReadString('\n')
		fmt.Println(str)
		/*if strings.Compare(string(str),"\n")==0 {
			fmt.Println("good bye")
			break
		}*/
		if len(str)!=0 {
			file.WriteString(string(str))
		} else{
			fmt.Println("good bye")
			break
		}
	}
	file.Close()

}
func main() {
  //bufioInput()
	writeFileWithbufioInput()
}
