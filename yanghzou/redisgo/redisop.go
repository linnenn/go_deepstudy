package redisgo
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	conf "yanghzou/utill"
)

var redisHandler redis.Conn
func init()  {
	url := fmt.Sprintf("%s:%d",conf.RedisHost,conf.RedisPort)
	var err error
	redisHandler,err = redis.Dial("tcp",url,redis.DialDatabase(1),redis.DialPassword(conf.RedisAuth))
	if err != nil {
		fmt.Println("connection redis error:",err)
		return
	}
	//redisHandler = conn
	//defer conn.Close()
	//stars := fmt.Sprintf("redis://%s:%d",conf.RedisHost,conf.RedisPort)
	//conn2,err := redis.DialURL(stars)
	//if err != nil {
	//	fmt.Println("connection redis error:",err)
	//	return
	//}
	//redisHandler = conn2
}
func PushList()  {
	//_,err := redisHandler.Do("lpush","fruits","apple","banana",300)
	//if err != nil {
	//	fmt.Println("push error ",err)
	//}
	res,_ := redis.String(redisHandler.Do("lpop","fruits"))
	fmt.Println(string(res))
}

func GetRecords()  {
	_,err := redisHandler.Do ("set","name","lei")
	if err != nil {
		fmt.Println("执行set失败",err)
	}
	res,err := redis.String(redisHandler.Do("get","name"))
	if err != nil {
		fmt.Println("执行get失败",err)
	}
	fmt.Println("the result:",string(res))
}


type Ints struct {
	A int
	B int
	C float32
}
func Mavut()  {
	i := Ints{1,2,1.2}
	args := redis.Args{1,2,3}
	args1 := args.AddFlat(i)
	fmt.Println(args,len(args))
	fmt.Println(args1,len(args1))
}

