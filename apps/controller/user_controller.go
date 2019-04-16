package user

import (
	"github.com/gin-gonic/gin"
	l "github.com/go_gin_sample/apps/configure/logger"
	"github.com/go_gin_sample/apps/service"
	"reflect"
	"strconv"
)

type Controller struct {
}
var logger  = l.GetLogger()

// ユーザー取得
func (pc *Controller) GetUser (c *gin.Context) {

	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		logger.Error(err)
		c.JSON(400, err)
		return
	}
	if id <= 0 {
		logger.Error(err)
		c.JSON(400, gin.H{"error": "id should be bigger than 0"})
		return
	}
	// データを処理する
	var service = user.Service{}
	result := service.GetById(id)

	if result == nil || reflect.ValueOf(result).IsNil() {
		logger.Error(err)
		c.JSON(404, gin.H{})
		return
	}
	c.JSON(200, result)
}