package controller

import (
	"github.com/gin-gonic/gin"
	l "github.com/go_gin_sample/apps/configure/logger"
	s "github.com/go_gin_sample/apps/domain/service"
	"github.com/go_gin_sample/apps/usecase"
	us "github.com/go_gin_sample/apps/usecase/user"
	"net/http"
	"strconv"
)
type UserController struct {
}

var logger  = l.GetLogger()

// ユーザー取得
func (pc *UserController) GetUser (c *gin.Context) {

	var getUserRequest us.GetUserRequest
	var commonResponse = usecase.CommonResponse{}
	if err := c.ShouldBindUri(&getUserRequest)

	// パラメータのチェック
	err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	id, _ := strconv.Atoi(getUserRequest.Id)
	// データを処理する
	var service = s.UserService{}
	response := service.GetById(id)

	c.JSON(200, commonResponse.CreateSuccessResponse(response))
}