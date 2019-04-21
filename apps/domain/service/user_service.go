package service

import (
	"errors"
	l "github.com/go_gin_sample/apps/configure/logger"
	"github.com/go_gin_sample/apps/domain/model"
	r "github.com/go_gin_sample/apps/infrastructure/repository"
)
var logger  = l.GetLogger()

type UserService struct {}

// IDを元にレコードを取得します
func (c *UserService) GetById(id uint64) (*model.User, error) {

	repo := &r.UserRepository{}
	user, err := repo.FindByUserId(id)

	return user, err
}


// Userの作成を行います
func (c *UserService) CreateUser(user *model.User) (*model.User, error) {

	repo := &r.UserRepository{}
	emailUser, emailErr := repo.FindByEmail(user.Email)
	if emailErr != nil {
		logger.Error("ユーザーの取得処理でエラーが発生しました。: ", emailErr)
		return nil, emailErr
	}

	if emailUser != nil {
		msg := "指定されたメールアドレスのユーザーは既に存在します。"
		logger.Warn("該当メールアドレスのユーザーは既に存在します。: ", emailUser.Email)
		return nil, errors.New(msg)
	}

	user.Status = model.UserStatusValid
	user, err := repo.Save(user)

	return user, err
}

// Userの更新を行います。
func (c *UserService) UpdateUser(paramUser *model.User) (*model.User, error) {

	repo := &r.UserRepository{}
	user, existsErr := repo.FindByUserId(paramUser.Id)
	if existsErr != nil {
		logger.Error("ユーザーの取得処理でエラーが発生しました。: ", existsErr)
		return nil, existsErr
	}

	if user == nil {
		msg := "指定されたユーザーが存在しません。"
		logger.Warn("指定されたユーザーが存在しません。ID: ", paramUser.Id)
		return nil, errors.New(msg)
	}

	// 同一メールアドレスのチェック
	emailUser, emailErr := repo.FindByEmail(paramUser.Email)
	if emailErr != nil {
		logger.Error("ユーザーの取得処理でエラーが発生しました。: ", emailErr)
		return nil, emailErr
	}

	if emailUser != nil && emailUser.Id != user.Id {
		msg := "指定されたメールアドレスのユーザーは既に存在します。"
		logger.Warn("該当メールアドレスのユーザーは既に存在します。: ", emailUser.Email)
		return nil, errors.New(msg)
	}

	// 値のコピー
	user.ValueCopy(paramUser)
	user, err := repo.Update(user)

	return user, err
}

// Userの削除を行います。
func (c *UserService) DeleteUser(id uint64) (*model.User, error) {

	repo := &r.UserRepository{}
	user, existsErr := repo.FindByUserId(id)
	if existsErr != nil {
		logger.Error("ユーザーの取得処理でエラーが発生しました。:", existsErr)
		return nil, existsErr
	}

	if user == nil {
		msg := "指定されたユーザーが存在しません。"
		logger.Warn("指定されたユーザーが存在しません。ID: ", id)
		return nil, errors.New(msg)
	}

	user, err := repo.Delete(user)

	return user, err
}