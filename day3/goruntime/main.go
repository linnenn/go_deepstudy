package main

import "fmt"

func main() {
	achan := make(chan int)
	go func() {
		for i := 10; i > 0; i-- {
			fmt.Println("puts >>>>", i)
			achan <- i
		}
		achan <- 0
	}()
	for vals := range achan {
		fmt.Println("gets <<<<", vals)
		if vals == 0 {
			fmt.Println("good bye", vals)
			break
		}
	}
	// achan <- 0
	fmt.Println("over")
}
