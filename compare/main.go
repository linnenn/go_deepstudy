package main

import "fmt"
// T interface
type T interface{}
type string1 string
type string2 = string

func main() {
	var a T = "aaa"
	var b string1 = "aaa"
	var c string2 = "aaa"
	fmt.Printf("%+v,%+v,%+v \n", a, b, c)

	fmt.Println(a == b)
	fmt.Println(a == string(b))
	fmt.Println(a == c)
	fmt.Println(a == string(c))


}
