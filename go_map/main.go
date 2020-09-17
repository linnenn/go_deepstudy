package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// NUM 1111
const NUM int = 10

func main() {
	aa := make(map[string]string, 20)
	for i := 0; i < NUM; i++ {
		aa[strconv.Itoa(i)] = fmt.Sprintf("%d", i)
	}
	m_json, _ := json.Marshal(aa)
	fmt.Println(string(m_json))

	// 删除某个元素
	delete(aa, "1")
	fmt.Println(aa)

}
