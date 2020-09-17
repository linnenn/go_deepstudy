package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	aa := [...]int{1, 2, 3, 4, 5, 56, 6}
	// for inde := range aa {
	// 	// for inde, _ := range aa {
	// 	fmt.Println(inde)
	// }
	// for {
	// 死循环
	// }
	fmt.Printf("%+v,%p,%d,%d\n", aa, &aa, len(aa), cap(aa))
	fmt.Printf("%p\n", &aa[0])
	fmt.Printf("%p\n", &aa[1])
	fmt.Printf("%p\n", &aa[2])
	fmt.Printf("%p\n", &aa[3])
	aa1 := aa[:]
	fmt.Printf("%p\n", &aa1[0])
	fmt.Printf("%p\n", &aa1[1])

	// aa1 {*aa,len,cap}
	// aa1[0] = 20
	fmt.Printf("%#v,%p,%d,%d\n", aa1, &aa1, len(aa1), cap(aa1))

	aa2 := aa[:3]
	fmt.Printf("%p\n", &aa2[0])
	fmt.Printf("%p\n", &aa2[1])

	fmt.Printf("%#v,%p,%d,%d\n", aa2, &aa2, len(aa2), cap(aa2))
	loopMap()

	srts := "jijijiiij"
	res := Md5String(srts)
	fmt.Println(res)
	ts := newtime()
	fmt.Println(ts)
	timestamp := nowTimestamp()
	fmt.Println(timestamp)
}

// LOOP aa
const LOOP = 20

func loopMap() {
	mp := make(map[string]string, 20)
	for i := 0; i < LOOP; i++ {
		rand.Seed(int64(time.Now().Unix()))

		mp[strconv.Itoa(i)] = strconv.Itoa(rand.Intn(LOOP))
	}
	fmt.Println(mp)
	for ind, val := range mp {
		fmt.Println(ind, val)
	}

}

// Md5String string
func Md5String(strigs string) string {
	md5str := md5.New()
	md5str.Write([]byte(strigs))
	return hex.EncodeToString(md5str.Sum(nil))
}

func newtime() string {
	times := time.Now().Format("2006/01/02 15:04:05")
	return times
}
func nowTimestamp() int64 {
	times := time.Now().Unix()
	return times
}
