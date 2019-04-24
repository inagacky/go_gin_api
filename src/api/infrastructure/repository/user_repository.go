package repository

import (
	"github.com/inagacky/go_gin_api/src/api/configure/db"
	l "github.com/inagacky/go_gin_api/src/api/configure/logger"
	m "github.com/inagacky/go_gin_api/src/api/domain/model"
	"github.com/jinzhu/gorm"
	"time"
)
var logger  = l.GetLogger()

type UserRepository interface {
	FindByUserId(id uint64) (*m.User, error)
	FindByEmail(email string) (*m.User, error)
	Save(user *m.User) (*m.User, error)
	Update(user *m.User) (*m.User, error)
	Delete(user *m.User) (*m.User, error)
}

type userRepository struct {
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
	}
}


// IDを元にユーザーを取得します
func (c *userRepository) FindByUserId(id uint64) (*m.User, error) {

	var user = m.User{}
	user.Id = id
	db := db.GetDB()
	if err := db.First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Error("ユーザーの取得処理でエラーが発生しました: ", err)
			return nil, err
		} else {
			logger.Info("ユーザーが存在しません。ID: ", id)
			return nil, nil
		}
	}

	return &user, nil
}

// Emailを元にユーザーの取得を行います。
func (c *userRepository) FindByEmail(email string) (*m.User, error) {

	var user = m.User{}
	user.Email = email
	db := db.GetDB().Where("email = ?", email)

	// Emailを元にユーザー取得
	if err := db.First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Error("ユーザーの取得処理でエラーが発生しました: ", err)
			return nil, err
		} else {
			logger.Info("ユーザーが存在しません。: ", email)
			return nil, nil
		}
	}

	return &user, nil
}

// ユーザー情報を作成します
func (c *userRepository) Save(user *m.User) (*m.User, error) {

	db := db.GetDB()
	if err := db.Create(&user).Error; err != nil {
		logger.Error("ユーザーの作成処理でエラーが発生しました: ", err)
		return nil, err
	}

	return user, nil
}

// ユーザー情報を更新します
func (c *userRepository) Update(user *m.User) (*m.User, error) {

	db := db.GetDB()
	if err := db.Save(&user).Error; err != nil {
		logger.Error("ユーザーの更新処理でエラーが発生しました: ", err)
		return nil, err
	}

	return user, nil
}

// ユーザー情報を削除します
func (c *userRepository) Delete(user *m.User) (*m.User, error) {

	db := db.GetDB()
	// 論理削除
	user.Status = m.UserStatusInValid
	current := time.Now()
	user.DeletedAt = &current

	if err := db.Save(&user).Error; err != nil {
		logger.Error("ユーザーの削除処理でエラーが発生しました: ", err)
		return nil, err
	}

	return user, nil
}