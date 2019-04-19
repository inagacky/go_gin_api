package controller

import (
	"github.com/gin-gonic/gin"
	l "github.com/go_gin_sample/apps/configure/logger"
	s "github.com/go_gin_sample/apps/domain/service"
	"github.com/go_gin_sample/apps/usecase"
	user_usecase "github.com/go_gin_sample/apps/usecase/user"
	"net/http"
	"reflect"
	"strconv"
)
type UserController struct {
}

var logger  = l.GetLogger()

// ユーザー取得
func (pc *UserController) GetUser (c *gin.Context) {

	var getUserRequest user_usecase.GetUserRequest
	if err := c.ShouldBindUri(&getUserRequest)

	err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, usecase.CommonResponse{}.CreateValidateErrorResponse(err.Error()))
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