package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main()  {
	r := gin.New()
	r.Use(costTime())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "首页")
	})
	group := r.Group("/admin")
	group.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))
	group.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"header":"header",
		})
	})

	r.Run(":8081")
}

func costTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Since(start)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, end)
	}
}
