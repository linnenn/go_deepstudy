package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, work chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "processor job", job)
		// time.Sleep(time.Second * 1)
		work <- job * 2
	}
}
func main() {
	jobs := make(chan int, 20)
	work := make(chan int, 20)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, work)
	}
	for i := 1; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 9; i++ {
		<-work
	}
}
