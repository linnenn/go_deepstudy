package main

import "fmt"

func main() {
	// 不管怎么做切片，最终的结果都是不可以修改的！！！字符串类型是不可变！！
	var a = 010
	//a += 010
	fmt.Println(a)

}
