package main

import "fmt"

func main() {
	var aa int
	var bb []int
	fmt.Println(aa, bb)
	if bb == nil {
		fmt.Println("hh")
	}
	fmt.Printf("%#v,%#v ", aa, bb)
}
