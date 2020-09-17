package main

import "fmt"

func main() {
	// 注意切片的容量和长度的使用，理解：长度是可访问空间，容量是内存是否需要申请新空间的标示
	// 切片默认值是：nil，其他基础类型和nil比较会报错，不同类型比较也不会相等
	// 切片共用一个底层数组，修改一个其他的可能会改变
	aa := make([]int, 3,10)
	fmt.Printf("%v,%p,len:%d,cap:%d\n",aa,aa,len(aa),cap(aa))
	for i := 0; i < 20; i++ {
		aa = append(aa,i)
		fmt.Printf("%v,%p\n",aa,aa,len(aa),cap(aa))
	}
	fmt.Printf("%v,%p\n",aa,aa,len(aa),cap(aa))
}