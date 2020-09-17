package main

import "fmt"

func main() {
	tzda
	var aaa = 32
	fmt.Printf("int,%#v \n",aaa)
	fmt.Printf("string,%v \n",string(aaa))
	var a = [...]int{1,2,3,4,5,6}
	b := a
	b[1] = 12
	//值语义，必然是不同的内存空间，不会互相影响
	fmt.Printf("%p,%p\n",&b,&a)
	fmt.Println(b,a)

}

type User struct {
	name string
}
func f1()  {
	user1 := User{
		name: "yabglie",
	}
	user2 := user1
	fmt.Println(user1,user2)
	fmt.Printf("%p,%p \n",&user2,&user1)

}

type stu map[string]string
func f2()  {
	a := make(stu,20)
	a["a"] = "aa"
	a["b"] = "bb"
	b := a
	b["a"] = "ccc"
	fmt.Println(b,a)
	fmt.Printf("%p,%p\n",&b,&a)
}