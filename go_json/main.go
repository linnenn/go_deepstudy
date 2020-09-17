package main

import (
	"encoding/json"
	"fmt"
)

// User user
type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
	Addr string `json:"addr"`
}

func main() {
	user := &User{
		Name: "uang",
		Age:  30,
		Addr: "beijing",
	}
	ujson, _ := json.Marshal(user)
	fmt.Println(string(ujson))

}
