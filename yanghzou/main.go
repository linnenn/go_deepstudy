package main

import (
	"yanghzou/router"
	_ "yanghzou/utill"
)


func main() {
	engine := router.InitRouter()
	engine.Run()
}
