package mathutil

//可变参数的应用，对任意多个数求和,bix将函数的第一个字母大写，否则为为导出主体，不能供外部使用
func AnySum(y...int)(total int) {
	total=0
	for _,v :=range y{
		total +=v
	}
	return
}
//可变参数的应用，对任意多个数求和,必须将函数的第一个字母大写，否则为为导出主体，不能供外部使用
func AnySMul(y...int)(total int) {
	total=1
	for _,v :=range y{
		total *=v
	}
	return
}

//可变参数的应用，获取任意一组数子的最大值
func AnyMax(y...int)(max int)  {
	max=y[0]
	for _,v:=range y{

		if max < v{
			max = v
		}
	}

	return max
}
//可变参数的应用，获取任意一组数字的最小值
func AnyMin(y...int)(min int){
	min = y[0]
	for _,v:=range y{
		if min > v{
			min = v
		}
	}
	return min
}
func GetAbs(x int)(abs int){
	if x < 0 {
		abs = x * -1
	} else{
		abs = x
	}
	return abs
}

func IsOdd(x int) bool {
	if(x%2==0){
		return false
	}else{
		return true
	}
}

func IsEven(x int) bool {
	if(x%2==0){
		return true
	}else{
		return false
	}
}