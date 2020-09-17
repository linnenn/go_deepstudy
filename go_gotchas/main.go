package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	var a,b = 12,1.5
	//command-line-arguments
	//./main.go:7:15: invalid operation: a / b (mismatched types int and float64)
	fmt.Println(float64(a)/b)
	var res = 12/1.5
	fmt.Printf("%T,%v \n",res,res)

	fmt.Println("*******************")
	var ma map[string]int
	//panic: assignment to entry in nil map
	//goroutine 1 [running]:
	//ma["home"] = 10
	//长度，默认值都是零值
	fmt.Println(ma["home"])
	fmt.Println(len(ma))
	//可以判断是否存在
	res1,ok := ma["foreign"]
	fmt.Println(res1,ok)
	fmt.Println("7&&&&&&&&&&&&&&&&")
	//下面的包括了初始化和赋值两种方式

	var bb = map[string]int{"hi":2}
	fmt.Println(bb["hi"])

	var cc = "aaa\tbbb\tccc"
	//aaa\tbbb\tccc
	//aaa     bbb     ccc
	fmt.Println(cc)

	var timeout = 3
	// command-line-arguments
	//./main.go:39:37: missing argument to conversion to time.Duration: time.Duration()
	fmt.Println(time.Duration(timeout) * time.Microsecond)

	const time11111= 12
	fmt.Println(time11111 * time.Second)

	var strings = "We♥Go"
	fmt.Println(len(strings))
	fmt.Println(utf8.RuneCountInString(strings))
}
