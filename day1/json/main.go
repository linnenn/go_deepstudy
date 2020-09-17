package main

import (
	"encoding/json"
	"fmt"
)

// People ss
type People struct {
	Name string `json:name,int`
	Age  int    `json:age`
	Addr string `json:addr`
}

type peoples []People

func main() {
	pp := make([]People, 20)
	for i := 0; i < 20; i++ {
		pp[i] = People{
			Name: fmt.Sprintf("%d", i),
			Age:  20 + i,
			Addr: fmt.Sprintf("北京市%d", i),
		}
	}
	// fmt.Println(pp)
	json, _ := json.Marshal(pp)
	fmt.Println(string(json))

}
