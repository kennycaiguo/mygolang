package main

//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
import (
	"fmt"
	"reflect"
)
type newInt int

func reflectType(x interface{})  {
	fmt.Printf("the type of x is :%v\n",reflect.TypeOf(x))
}

func reflectTypekind(x interface{})  {
	fmt.Printf("the type of x is :%v\n,kind of x :%v",reflect.TypeOf(x).Name(),reflect.TypeOf(x).Kind())
}
func setValue(x interface{}){
	v:=reflect.ValueOf(x)
	if v.Elem().Kind()==reflect.Int64{
       v.Elem().SetInt(300)
	}
}

func main() {
	/*var x float32 = 3.2
	reflectType(x)*/
	/*var y int64 =100
	reflectTypekind(y)*/
	//var a newInt =30 //the type of x is :newInt ,kind of x :int
	var b int64 =30 //the type of x is :newInt ,kind of x :int

	//reflectTypekind(a)
	setValue(&b) //这里要传b的地址
	fmt.Printf("%d",b)
}
