package main

import "fmt"

func main() {
	var aa []int
	var bb []int
	// 下面的全部成功
	if aa == nil {
		fmt.Println("is nil")
	}
	if bb == nil {
		fmt.Println("is nil")
	}
	//invalid operation: aa == bb (slice can only be compared to nil)
	// 也就是说不能比较
	// if aa == bb {
	// 	fmt.Println("is equal")
	// }
	var cc map[int]string
	// 下面的成功
	if cc == nil {
		fmt.Println("is nil")
	}
	var dd map[int]string
	//  invalid operation: cc == dd (map can only be compared to nil)
	// if cc == dd {
	// 	fmt.Println("is equal")
	// }
	// 不同类型的nil不可以比较
	// invalid operation: bb == dd (mismatched types []int and map[int]string)
	if bb == dd {
		fmt.Println("不通类型的nil比较")

	}
}
