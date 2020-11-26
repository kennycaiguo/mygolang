package main

import "fmt"

func main() {
	/*m:=map[string]int{"age":20,"height":175} //创建map对象并且赋值
	println(m) //0xc00003be78
    fmt.Println(m) //map[age:20 height:175] 所以要用fmt.Println
    fmt.Println(m["age"]) //20
	for s := range m {
		fmt.Println(s)    //age "\n" height
		fmt.Println(m[s]) //20 "\n" 175
	}*/
    var map1 map[string]int
    map1 = make(map[string]int,1) //可以使用make函数
    fmt.Println(len(map1))
    map1["money"]=100
    map1["debt"]=20  //key 和 value可以动态赋值，map可以自动扩容，map是无序的
    fmt.Println(map1) //map[debt:20 money:100]
    //value,ok :=map1["pussy"] //测试查询一个key，不存在就提示没有，存在就打印,结果：查无此键
    value,ok :=map1["debt"] //测试查询一个key，不存在就提示没有，存在就打印value，结果:20
    if !ok{
    	fmt.Println("查无此键")
	}else{
		fmt.Println(value)
	}
	map1["pussy"]=12
	fmt.Println(map1)
	delete(map1,"pussy")//map删除元素
	fmt.Println(map1)
}
