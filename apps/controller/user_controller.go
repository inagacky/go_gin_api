package controller

import (
	"github.com/gin-gonic/gin"
	l "github.com/go_gin_sample/apps/configure/logger"
	s "github.com/go_gin_sample/apps/domain/service"
	usecase "github.com/go_gin_sample/apps/usecase/user"
	"net/http"
	"reflect"
	"strconv"
)
type UserController struct {
}

var logger  = l.GetLogger()

// ユーザー取得
func (pc *UserController) GetUser (c *gin.Context) {

	var getUserRequest usecase.GetUserRequest
	if err := c.ShouldBindUri(&getUserRequest)

	err != nil {
/*
		errors := err.(*validator.StructErrors)
		log.Println("Struct:", errors.Struct)
		for k, v := range errors.Errors {
			log.Println("Key:", k)
			log.Println("Field:", v.Field)
			log.Println("Param:", v.Param)
			log.Println("Tag:", v.Tag)
			log.Println("Kind", v.Kind)
			log.Println("Type:", v.Type)
			log.Println("Value", v.Value)
			log.Println("==========")
		}
		*/
		logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest", "error": err.Error()})
		return
	}
	// パラメータ取得
	n := c.Param("id")
	id, err := strconv.Atoi(n)

	// データを処理する
	var service = s.UserService{}
	result := service.GetById(id)

	if result == nil || reflect.ValueOf(result).IsNil() {
		logger.Error(err)
		c.JSON(404, gin.H{})
		return
	}
	c.JSON(200, result)
}