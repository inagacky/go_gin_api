package controller

import (
	"github.com/gin-gonic/gin"
	l "github.com/inagacky/go_gin_api/src/api/configure/logger"
	s "github.com/inagacky/go_gin_api/src/api/domain/service"
	"github.com/inagacky/go_gin_api/src/api/usecase"
	us "github.com/inagacky/go_gin_api/src/api/usecase/user"
	"net/http"
	"strconv"
)

type UserController interface {
	GetUser (c *gin.Context)
	CreateUser (c *gin.Context)
	UpdateUser (c *gin.Context)
	DeleteUser (c *gin.Context)
}

func NewUserController(userService s.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

type userController struct {
	userService s.UserService
}

// ユーザー取得API
func (co *userController) GetUser (c *gin.Context) {

	var getUserRequest us.GetUserRequest
	commonResponse := &usecase.CommonResponse{}
	getUserRequest.Id = c.Param("id")

	// パラメータのチェック
	if err := c.Bind(&getUserRequest); err != nil {
		l.GetLogger().Error(err.Error())

		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// int64への変換
	id, _ := strconv.ParseUint(getUserRequest.Id, 10 ,64)
	// ユーザー取得
	user, err := co.userService.GetById(id)
	if err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(us.GetUserResponse{User:user}))
}


// ユーザー作成API
func (co *userController) CreateUser (c *gin.Context) {

	var createUserRequest us.CreateUserRequest

	commonResponse := &usecase.CommonResponse{}
	// パラメータのチェック
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}
	// ユーザー作成
	user, err := co.userService.CreateUser(createUserRequest.ConvertUserModel())
	if err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(us.CreateUserResponse{User:user}))
}

// ユーザー情報更新API
func (co *userController) UpdateUser (c *gin.Context) {

	var updateUserRequest us.UpdateUserRequest
	commonResponse := &usecase.CommonResponse{}

	updateUserRequest.Id = c.Param("id")
	// パラメータのチェック
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// ユーザー更新
	user, err := co.userService.UpdateUser(updateUserRequest.ConvertUserModel())
	if err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(us.CreateUserResponse{User:user}))
}

// ユーザー削除API
func (co *userController) DeleteUser (c *gin.Context) {

	var deleteUserRequest us.DeleteUserRequest
	commonResponse := &usecase.CommonResponse{}
	deleteUserRequest.Id = c.Param("id")

	// パラメータのチェック
	if err := c.Bind(&deleteUserRequest); err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, commonResponse.CreateValidateErrorResponse(err.Error()))
		return
	}

	// int64への変換
	id, _ := strconv.ParseUint(deleteUserRequest.Id, 10, 64)
	// ユーザー削除
	user, err := co.userService.DeleteUser(id)
	if err != nil {
		l.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, commonResponse.CreateSQLErrorResponse(err.Error()))
		return
	}

	c.JSON(200, commonResponse.CreateSuccessResponse(us.DeleteUserResponse{User: user}))
}