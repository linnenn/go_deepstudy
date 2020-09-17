package main

import "fmt"

func main() {
	aa := []int{1, 2, 3, 4, 5}
	fmt.Printf("before:%p,%T,%#v\n", aa, aa, aa)
	sliceInfo(aa)
	fmt.Printf("after:%p,%T,%#v\n", aa, aa, aa)
}

// 看上去是引用传值，其实是因为修改内部的指针，通过指针当然可以修改底层数组
// 传递过来的是切片的不变部分，比如：指针，len，cap，指针指向的内存这一部分是可以改变的
// 所以说在函数内部修改参数是可以的，但是只要牵涉到切片的可变部分，就不再是原来的切片，就变成了新的切片
func sliceInfo(ss []int) {
	fmt.Printf("In1:%p,%T,%#v\n", ss, ss, ss)
	// 下面的操作原数组会改变
	for i := 0; i < len(ss); i++ {
		ss[i] = 2 * i
	}
	// 下面的会生成新的指针位置，不再是原来的切片，修改不再对外面的变量做修改
	ss = append(ss, 1)
	// ss[1] = 1111111
	fmt.Printf("In2:%p,%T,%#v\n", ss, ss, ss)
}
