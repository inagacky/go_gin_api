package usecase

import (
	"errors"
	l "github.com/inagacky/go_gin_api/src/api/configure/logger"
	"github.com/inagacky/go_gin_api/src/api/domain/entity"
	r "github.com/inagacky/go_gin_api/src/api/infrastructure/repository"
	"strconv"
)

type UserService interface {
	GetById(id uint64) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(paramUser *entity.User) (*entity.User, error)
	DeleteUser(id uint64) (*entity.User, error)
}

func NewUserService(userRepository r.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

type userService struct {
	userRepository r.UserRepository
}

// IDを元にレコードを取得します
func (c *userService) GetById(id uint64) (*entity.User, error) {

	user, err := c.userRepository.FindByUserId(id)

	return user, err
}


// Userの作成を行います
func (c *userService) CreateUser(user *entity.User) (*entity.User, error) {

	emailUser, emailErr := c.userRepository.FindByEmail(user.Email)
	if emailErr != nil {
		l.GetLogger().Error("ユーザーの取得処理でエラーが発生しました。: " + emailErr.Error())
		return nil, emailErr
	}

	if emailUser != nil {
		msg := "指定されたメールアドレスのユーザーは既に存在します。"
		l.GetLogger().Warn("該当メールアドレスのユーザーは既に存在します。: " + emailUser.Email)
		return nil, errors.New(msg)
	}

	user.Status = entity.UserStatusValid
	user, err := c.userRepository.Save(user)

	return user, err
}

// Userの更新を行います。
func (c *userService) UpdateUser(paramUser *entity.User) (*entity.User, error) {

	user, existsErr := c.userRepository.FindByUserId(paramUser.Id)
	if existsErr != nil {
		l.GetLogger().Error("ユーザーの取得処理でエラーが発生しました。: " + existsErr.Error())
		return nil, existsErr
	}

	if user == nil {
		msg := "指定されたユーザーが存在しません。"
		l.GetLogger().Warn("指定されたユーザーが存在しません。ID: " + strconv.FormatUint(paramUser.Id, 10))
		return nil, errors.New(msg)
	}

	// 同一メールアドレスのチェック
	emailUser, emailErr := c.userRepository.FindByEmail(paramUser.Email)
	if emailErr != nil {
		l.GetLogger().Error("ユーザーの取得処理でエラーが発生しました。: " + emailErr.Error())
		return nil, emailErr
	}

	if emailUser != nil && emailUser.Id != user.Id {
		msg := "指定されたメールアドレスのユーザーは既に存在します。"
		l.GetLogger().Warn("該当メールアドレスのユーザーは既に存在します。: " + emailUser.Email)
		return nil, errors.New(msg)
	}

	// 値のコピー
	user.ValueCopy(paramUser)
	user, err := c.userRepository.Update(user)

	return user, err
}

// Userの削除を行います。
func (c *userService) DeleteUser(id uint64) (*entity.User, error) {

	user, existsErr := c.userRepository.FindByUserId(id)
	if existsErr != nil {
		l.GetLogger().Error("ユーザーの取得処理でエラーが発生しました。:" + existsErr.Error())
		return nil, existsErr
	}

	if user == nil {
		msg := "指定されたユーザーが存在しません。"
		l.GetLogger().Warn("指定されたユーザーが存在しません。ID: "+ strconv.FormatUint(id, 10))
		return nil, errors.New(msg)
	}

	user, err := c.userRepository.Delete(user)

	return user, err
}
