package main

import "fmt"

func main() {
	var a []int = []int{1: 0}
	fmt.Println(a)
	var b []int
	// b[0] = 1
	fmt.Println(b, len(b), cap(b))
	fmt.Printf("%+v,%p \n", b, b)
	c := []int{}
	// c[0] = 10
	fmt.Println(c, len(c), cap(c))
	fmt.Printf("%+v,%p\n", c, c)

}
