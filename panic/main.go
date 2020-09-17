package main

import "fmt"

func main()  {
	defer func() {
		fmt.Println("aaa")
	}()
	defer func() {
		fmt.Println("bbbb")
	}()
	defer func() {
		fmt.Println("cccc")
	}()
	panic("error")
}