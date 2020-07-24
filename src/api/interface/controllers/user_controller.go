package controllers

import (
	"github.com/gin-gonic/gin"
	l "github.com/inagacky/go_gin_api/src/api/configure/logger"
	"github.com/inagacky/go_gin_api/src/api/interface/http"
	"github.com/inagacky/go_gin_api/src/api/usecase"
	net "net/http"
	"strconv"
)

type UserController interface {
	GetUser (c *gin.Context)
	CreateUser (c *gin.Context)
	UpdateUser (c *gin.Context)
	DeleteUser (c *gin.Context)
}

func NewUserController(userService usecase.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

type userController struct {
	userService usecase.UserService
}

// ユーザー取得API
func (co *userController) GetUser (c *gin.Context) {

	var getUserRequest http.GetUserRequest
	commonResponse := &http.CommonResponse{}
	getUserRequest.Id = c.Param("id")

	// パラメータのチェック
	if err := c.Bind(&getUserRequest); err != nil {
		l.GetLogger().Error(err.Error())

		c.JSON(net.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// int64への変換
	id, _ := strconv.ParseUint(getUserRequest.Id, 10 ,64)
	// ユーザー取得
	user, err := co.userService.GetById(id)
	if err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(net.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(http.GetUserResponse{User: user}))
}


// ユーザー作成API
func (co *userController) CreateUser (c *gin.Context) {

	var createUserRequest http.CreateUserRequest

	commonResponse := &http.CommonResponse{}
	// パラメータのチェック
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(net.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}
	// ユーザー作成
	user, err := co.userService.CreateUser(createUserRequest.ConvertUserModel())
	if err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(net.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(http.CreateUserResponse{User: user}))
}

// ユーザー情報更新API
func (co *userController) UpdateUser (c *gin.Context) {

	var updateUserRequest http.UpdateUserRequest
	commonResponse := &http.CommonResponse{}

	updateUserRequest.Id = c.Param("id")
	// パラメータのチェック
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(net.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// ユーザー更新
	user, err := co.userService.UpdateUser(updateUserRequest.ConvertUserModel())
	if err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(net.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(http.CreateUserResponse{User: user}))
}

// ユーザー削除API
func (co *userController) DeleteUser (c *gin.Context) {

	var deleteUserRequest http.DeleteUserRequest
	commonResponse := &http.CommonResponse{}
	deleteUserRequest.Id = c.Param("id")

	// パラメータのチェック
	if err := c.Bind(&deleteUserRequest); err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(net.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// int64への変換
	id, _ := strconv.ParseUint(deleteUserRequest.Id, 10, 64)
	// ユーザー削除
	user, err := co.userService.DeleteUser(id)
	if err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(net.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(http.DeleteUserResponse{User: user}))
}
