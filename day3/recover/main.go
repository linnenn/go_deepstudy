package main

import (
	"fmt"
	"math/rand"
)

type divider struct {
	first int
	second int
}
func main()  {
	var divd = make([]*divider,10,30)
	for i := 0;i<10;i++ {
		divd[i] = &divider{
			first: rand.Intn(10)+20,
			second: 0,
		}
	}
	for _,record := range divd{
		if result,err := record.f1(); err != nil {
			fmt.Println(result,err)
		}
	}
}

func (d *divider) f1() (a int, errs error) {
	//if d.second == 0 {
	//	return 0, fmt.Errorf("被除数second=%v不能作为除数",d.second)
	//}
	defer func() {
		if err := recover();err != nil {
			fmt.Println(err)
			a = -1
			//errs = errors.New("被除数不能作为除数")
		}
	}()
	return d.first / d.second,nil
}