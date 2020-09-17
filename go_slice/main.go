package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	// 追加单个，追加多个，追加数组切片
	// a = append(a, 1) //[1 2 3 4 5 6 1] 7 12
	// a = append(a, 1, 2, 3, 4) //[1 2 3 4 5 6 1 2 3 4] 10 12
	// a = append(a, []int{3, 4, 5}...) //[1 2 3 4 5 6 3 4 5] 9 12
	fmt.Println(a, len(a), cap(a))
	// 切片的长度和容量的理解
	// b := a[:] //[1 2 3 4 5 6] 6 6
	// b := a[:7] //[1 2 3 4 5 6] 6 6//下标不能超过长度， slice bounds out of range [:7] with capacity 6
	// b := a[:5] //[1 2 3 4 5] 5 6
	// b := a[1:] //[2 3 4 5 6] 5 5
	// b := a[5:] //[6] 1 1
	//切片复制，copy,注意容量和复制的关系
	// b := []int{}
	// copy(b, a) //[] 0 0
	// b := make([]int, 1)
	// copy(b, a) //[1] 1 1
	// b := make([]int, 10)
	// copy(b, a) //[1 2 3 4 5 6 0 0 0 0] 10 10
	// 删除切片，删除切片并没有定义语法糖，需要使用切片来做处理
	// b := make([]int, 10)
	copy(b, a) //[1 2 3 4 5 6 0 0 0 0] 10 10
	// 删除第二个,删除第几个对切片做处理
	b = append(b[0:1], b[2:]...) //[1 3 4 5 6 0 0 0 0] 9 10
	fmt.Println(b, len(b), cap(b))
}
