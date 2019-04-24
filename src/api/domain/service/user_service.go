package service

import (
	"errors"
	l "github.com/inagacky/go_gin_api/src/api/configure/logger"
	"github.com/inagacky/go_gin_api/src/api/domain/model"
	r "github.com/inagacky/go_gin_api/src/api/infrastructure/repository"
)
var logger  = l.GetLogger()

type UserService interface {
	GetById(id uint64) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(paramUser *model.User) (*model.User, error)
	DeleteUser(id uint64) (*model.User, error)
}

type userService struct {
	userRepository r.UserRepository
}

func NewUserService(userRepository r.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// IDを元にレコードを取得します
func (c *userService) GetById(id uint64) (*model.User, error) {

	user, err := c.userRepository.FindByUserId(id)

	return user, err
}


// Userの作成を行います
func (c *userService) CreateUser(user *model.User) (*model.User, error) {

	emailUser, emailErr := c.userRepository.FindByEmail(user.Email)
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
	user, err := c.userRepository.Save(user)

	return user, err
}

// Userの更新を行います。
func (c *userService) UpdateUser(paramUser *model.User) (*model.User, error) {

	user, existsErr := c.userRepository.FindByUserId(paramUser.Id)
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
	emailUser, emailErr := c.userRepository.FindByEmail(paramUser.Email)
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
	user, err := c.userRepository.Update(user)

	return user, err
}

// Userの削除を行います。
func (c *userService) DeleteUser(id uint64) (*model.User, error) {

	user, existsErr := c.userRepository.FindByUserId(id)
	if existsErr != nil {
		logger.Error("ユーザーの取得処理でエラーが発生しました。:", existsErr)
		return nil, existsErr
	}

	if user == nil {
		msg := "指定されたユーザーが存在しません。"
		logger.Warn("指定されたユーザーが存在しません。ID: ", id)
		return nil, errors.New(msg)
	}

	user, err := c.userRepository.Delete(user)

	return user, err
}