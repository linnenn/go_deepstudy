package main

import "fmt"

type user struct {
	name string
	age  int
	addr string
}

func main() {
	// new 初始化返回指针
	user1 := new(user)
	user1.name = "yanglei"
	user1.age = 32
	fmt.Println(user 1)
	// 直接初始化
	user2 := user{name: "yanglei", age: 12}
	fmt.Println(user2)
	// 直接初始化，返回指针，这是一个语法糖
	user3 := &user{name: "hhh", age: 32}
	fmt.Println(user3)
	// 直接初始化，顺序给出字段,并且一个不能少：too few values in user literal
	// 包括指针类型和实际的实例化类型
	// user4 := user{"yanggggggg", 20, "河南"}
	user4 := &user{"yanggggggg", 20, "河南"}
	fmt.Println(user4)

	// 数组json格式化，需要使用当前包中的变量，所以必须外部可见，需要首字母大写

}
