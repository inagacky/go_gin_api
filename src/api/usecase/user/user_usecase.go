package usecase

import (
	"github.com/inagacky/go_gin_sample/src/api/domain/model"
	"strconv"
)

// ユーザー取得のリクエスト
type GetUserRequest struct {
	Id string `uri:"id" binding:"number,required,min=1"`
}

// ユーザー取得のレスポンス
type GetUserResponse struct {
	User *model.User `json:"user"`
}

// ユーザー作成のリクエスト
type CreateUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

// ユーザー作成のレスポンス
type CreateUserResponse struct {
	User *model.User `json:"user"`
}

// ユーザー更新のリクエスト
type UpdateUserRequest struct {
	Id string `uri:"id" binding:"number,required,min=1"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

// ユーザー更新のレスポンス
type UpdateUserResponse struct {
	User *model.User `json:"user"`
}

// ユーザー削除のリクエスト
type DeleteUserRequest struct {
	Id string `uri:"id" binding:"number,required,min=1"`
}

// ユーザー削除のレスポンス
type DeleteUserResponse struct {
	User *model.User `json:"user"`
}



func (request *CreateUserRequest) ConvertUserModel () *model.User {

	user := &model.User{}
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email

	return user
}

func (request *UpdateUserRequest) ConvertUserModel () *model.User {

	user := &model.User{}
	user.Id, _ = strconv.ParseUint(request.Id, 10 ,64)
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email

	return user
}