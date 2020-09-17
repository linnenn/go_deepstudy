package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// User user
type User struct {
	UserName string `json:"name"`
	UserAge  int    `json:"age"`
	UserAddr string `json:"addr"`
}

func main() {
	uss := &User{
		UserName: "lei",
		UserAge:  44,
		UserAddr: "beijing",
	}
	fmt.Println(uss)
	ustring := Md5sign(*uss)
	fmt.Println(ustring)
}

// Md5sign strg
func Md5sign(uss User) string {
	ms := md5.New()
	ms.Write([]byte(uss.UserName + uss.UserAddr))
	return hex.EncodeToString(ms.Sum(nil))
}
