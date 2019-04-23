package controller

import (
	"github.com/gin-gonic/gin"
	l "github.com/inagacky/go_gin_sample/src/api/configure/logger"
	s "github.com/inagacky/go_gin_sample/src/api/domain/service"
	"github.com/inagacky/go_gin_sample/src/api/usecase"
	us "github.com/inagacky/go_gin_sample/src/api/usecase/user"
	"net/http"
	"strconv"
)
type UserController struct {
}

var logger  = l.GetLogger()

// ユーザー取得API
func (pc *UserController) GetUser (c *gin.Context) {

	var getUserRequest us.GetUserRequest
	commonResponse := &usecase.CommonResponse{}
	getUserRequest.Id = c.Param("id")

	// パラメータのチェック
	if err := c.Bind(&getUserRequest); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// int64への変換
	id, _ := strconv.ParseUint(getUserRequest.Id, 10 ,64)
	// ユーザー取得
	service := &s.UserService{}
	user, err := service.GetById(id)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(us.GetUserResponse{User:user}))
}


// ユーザー作成API
func (pc *UserController) CreateUser (c *gin.Context) {

	var createUserRequest us.CreateUserRequest

	commonResponse := &usecase.CommonResponse{}
	// パラメータのチェック
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}
	// ユーザー作成
	service := &s.UserService{}
	user, err := service.CreateUser(createUserRequest.ConvertUserModel())
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(us.CreateUserResponse{User:user}))
}

// ユーザー情報更新API
func (pc *UserController) UpdateUser (c *gin.Context) {

	var updateUserRequest us.UpdateUserRequest
	commonResponse := &usecase.CommonResponse{}

	updateUserRequest.Id = c.Param("id")
	// パラメータのチェック
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// ユーザー更新
	service := &s.UserService{}
	user, err := service.UpdateUser(updateUserRequest.ConvertUserModel())
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(us.CreateUserResponse{User:user}))
}

// ユーザー削除API
func (pc *UserController) DeleteUser (c *gin.Context) {

	var deleteUserRequest us.DeleteUserRequest
	commonResponse := &usecase.CommonResponse{}
	deleteUserRequest.Id = c.Param("id")

	// パラメータのチェック
	if err := c.Bind(&deleteUserRequest); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// int64への変換
	id, _ := strconv.ParseUint(deleteUserRequest.Id, 10, 64)
	// ユーザー削除
	service := &s.UserService{}
	user, err := service.DeleteUser(id)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(us.DeleteUserResponse{User: user}))
}