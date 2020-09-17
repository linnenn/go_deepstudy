package main

import (
	"fmt"
)

type Map map[string][]int

func main() {
	fmt.Println('z','Z')
	getChineseNum()
	//f,_ := os.Create("out_trace.out")
	//defer f.Close()
	//err := trace.Start(f)
	//if err != nil {
	//	panic(err)
	//}
	//defer trace.Stop()
	//
	//
	//go printInfo()
	//time.Sleep(time.Nanosecond * 100)
	//fmt.Println("prit info")
}

func getChineseNum() {
	strs := "hello沙河小王子"
	for _,va := range  strs{
		if va <= 122 {
			continue
		}
		fmt.Println(string(va))
	}
}
func printInfo() {
	fmt.Println("printInfo function ")
}

func findIndex()  {
	var a = [...]int{1, 3, 5, 7, 8}
	var nn = make([]int, 0)
	for i, v := range a {
		for ii, vv := range a {
			if v+vv == 8 && i <= ii {
				nn = append(nn,i,ii)
			}
		}
	}
	fmt.Println(nn)
}
