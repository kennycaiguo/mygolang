package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {

	count:=0
	for i:=0;i< len(users);i++{
		//str:=[]rune(users[i])
		str:=strings.Split(users[i],"")
		for _,c :=range str{
			 //fmt.Println(c)
			if(c=="e"||c=="E"){
				count++
				coins--
			}else if(c=="i"||c=="I"){
				count+=2
				coins-=2
			}else if(c=="o"||c=="O"){
				count+=3
				coins-=3
			}else if(c=="u"||c=="U"){
				count+=4
				coins-=4
			}
		}
		distribution[users[i]]=count
		count=0
	}
	//distribution["coins"]=coins
	fmt.Println(distribution)
	return coins
}
func main() {
   left:=dispatchCoin()
   fmt.Printf("金币分完了还剩下%d个",left);
}
