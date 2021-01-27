# mygolang
# <a href="https://www.liwenzhou.com/">go语言李文周老师的博客</a>
# <a href="https://www.cnblogs.com/liang1101/p/6741262.html">GoLang中 json、map、struct 之间的相互转化</a>
# <a href="https://www.runoob.com/go/go-recursion.html">golang中函数的递归调用</a>
# <a href="http://sinhub.cn/2019/12/use-golang-connect2oracle-on-windows/">golang 操作Oracle数据库</a>
# <a href="https://article.itxueyuan.com/e9DL7">golang 操作Oracle数据库2</a>
# <a href="https://juejin.im/post/6844904122143997965">golang高性能日志库zap配置示例</a>
# <a href="https://juejin.im/post/6844904099788357639">关于收集，标准化和集中化处理Golang日志的一些建议</a>
# <a href="https://segmentfault.com/a/1190000020992460">windows安装配置redis server</a>
# golang 读取文本文件内容
func readTextFile(path string) {
	f,_:=os.OpenFile(path,os.O_RDONLY,0666)
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

# go语言在文本文件指定的位置插入内容
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

# golang 实现求整数的平方根
package main

import "fmt"

func mySqrt1(x int) int {
	res:= x
    //牛顿法求平方根
    for res*res > x {
        res = (res + x/res) / 2
    }
    return res
}

func main() {
	a:=625
	fmt.Printf("%d的平方根是%d",a,mySqrt1(a))
}

# 结构体的匿名字段的使用
package main

import "fmt"

type person struct{
	string
	int
}
func main() {
	p1:=person{
		"kenny",
		35,
	}
	
   fmt.Printf("my name is %s,i am %d!",p1.string,p1.int)
	
}

# golang 结构体嵌套实例，写法1
package main

import "fmt"

type Address struct{
	province string
	city     string
}
type person struct{
	 name string
	 age   int
	 address Address
}

func main() {
	addr:=Address{
	    "广东省",
	    "东莞市",
	}
	p1:=person{
		name:"kenny",
		age:35,
		address:addr,
		
	}
	
   fmt.Printf("my name is %s,i am %d!\n",p1.name,p1.age)
	fmt.Printf("my address is %s,%s\n",p1.address.province,p1.address.city)
}
结果：
my name is kenny,i am 35!
my address is 广东省,东莞市

## 结构体嵌套实例，写法2
package main

import "fmt"

type Address struct{
	province string
	city     string
}
type person struct{
	 name string
	 age   int
	 address Address
}

func main() {
	 
	p1:=person{
		name:"kenny",
		age:35,
		address:Address{
	    "广东省",
		"东莞市",
	  },
		
	}
	
   fmt.Printf("my name is %s,i am %d!\n",p1.name,p1.age)
	fmt.Printf("my address is %s,%s\n",p1.address.province,p1.address.city)
}
my name is kenny,i am 35!
my address is 广东省,东莞市

# 匿名嵌套结构体实例
package main

import "fmt"

type Address struct{
	province string
	city     string
}
type person struct{
	 name string
	 age   int
     Address
}

func main() {
	 
	p1:=person{
		name:"kenny",
		age:35,
		Address:Address{
	    "广东省",
		"东莞市",
	  },
		
	}
	
   fmt.Printf("my name is %s,i am %d!\n",p1.name,p1.age) //使用匿名嵌套结构体的好处：可以简化操作
   fmt.Printf("my address is %s,%s\n",p1.province,p1.city)
}

结果：
my name is kenny,i am 35!
my address is 广东省,东莞市

