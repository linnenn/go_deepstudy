package main

import (
	"fmt"
	"strings"
	"time"
)

//作用域导致显示的数据只能在定义的范围内显示
func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret := 10
		defer func() {
			ret = ret + 1
			fmt.Println("in",ret)
		}()
		fmt.Println("out",ret)
	}
	return
}

// command-line-arguments
//./main.go:23:3: ret is shadowed during return

func A() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}

func B() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}
//2，3 随机输出，因为都是可以执行的，但是最终的结果会因为不同的方法的休眠时间导致获得最终结果的方法
func main() {
	//下面的去除是排列与组合删除
	fmt.Println(strings.TrimRight("ABABAC", "BCA"))
	fmt.Println(strings.TrimRight("abba", "ba"))
	fmt.Println(strings.TrimRight("abcdaaaaa", "abcd"))
	//只删除相同的后缀，也有前缀
	fmt.Println(strings.TrimSuffix("abcddcba", "dcba"))
	min(123,456)
	//ch := make(
	//chan int, 1)
	//go func() {
	//	select {
	//		case ch <- A():
	//		case ch <- B():
	//		default:
	//			ch <- 3
	//	}
	//}()
	//fmt.Println(<-ch)
}
//default是关键字，不能定义下面的函数
//func C(val, default string) string {
//	if val == "" {
//		return default
//}
//return val
//}

func min(a int, b uint) {
	var min = 0
	//因为copy返回的是拷贝之后的最小值的大小
	min = copy(make([]struct{},a),make([]struct{},b))
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}