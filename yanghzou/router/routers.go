package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanghzou/service"
)

//初始化返回路由
func InitRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"ping":"pong",
		})
	})
	user := engine.Group("/user")
	{
		//增加

		.POST("",service.UserAdd)
		//更新
		user.PUT("/:id",service.UserUpdate)
		//删除
		user.DELETE("/:id",service.UserDelete)
		//获取个人信息
		user.GET("/:id",service.UserInfo)
		//查询信息
		user.GET("",service.UserIndex)
	}
	return engine
}
