package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanghzou/model"
)

func UserIndex(c *gin.Context) {
	page := c.DefaultQuery("page","10")
	records := model.GetRecords()
	c.JSON(http.StatusOK,gin.H{
		"data":records,
		"page":page,
	})
}
func UserAdd(c *gin.Context)  {

}

func UserUpdate(c *gin.Context)  {
	
}

func UserDelete(c *gin.Context)  {
	
}

type BindUser struct {
	Uid int `json:"uid" binding:"required"`
}
func UserInfo(c *gin.Context)  {
	var binder BindUser
	if err := c.ShouldBindUri(&binder);err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"data":model.GetUserRecord(binder.Uid),
	})
}