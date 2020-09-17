package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type user struct {
	name string
}

type Ints struct {
	A int
	B int
	C float32
}
func main()  {
	i := Ints{1,2,1.2}
	args := redis.Args{1,2,3}
	args1 := args.AddFlat(i)
	fmt.Println(args,len(args))
	fmt.Println(args1,len(args1))
}
