package main

import "fmt"

func main() {
	aa := make(map[int]string, 20)
	fmt.Printf("1:%p,%T,%#v\n", aa, aa, aa)
	aa[1] = "1a"
	aa[2] = "2a"
	aa[3] = "3a"
	maps(aa)
	fmt.Printf("2:%p,%T,%#v\n", aa, aa, aa)

}

func maps(mm map[int]string) {
	fmt.Printf("11:%p,%T,%#v\n", mm, mm, mm)
	mm[1] = "11111"
	fmt.Printf("12:%p,%T,%#v\n", mm, mm, mm)
}
