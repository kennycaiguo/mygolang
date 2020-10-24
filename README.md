# mygolang
# <a href="https://www.liwenzhou.com/">go语言李文周老师的博客</a>
# <a href="https://www.cnblogs.com/liang1101/p/6741262.html">GoLang中 json、map、struct 之间的相互转化</a>
# <a href="https://www.runoob.com/go/go-recursion.html">golang中函数的递归调用</a>
# <a href="http://sinhub.cn/2019/12/use-golang-connect2oracle-on-windows/">golang 操作Oracle数据库</a>
# <a href="https://article.itxueyuan.com/e9DL7">golang 操作Oracle数据库2</a>

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

#结构体的匿名字段的使用
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
