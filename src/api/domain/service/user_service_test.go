package service


import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/inagacky/go_gin_api/src/api/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


var timeNow = time.Now()

// 正常値を返却するモック
type UserRepositorySuccessMock struct {
}

func (m *UserRepositorySuccessMock) FindByUserId(id uint64) (*model.User, error) {

	return createTestUser(), nil
}

func (m *UserRepositorySuccessMock) FindByEmail(email string) (*model.User, error) {

	return nil, nil
}

func (m *UserRepositorySuccessMock) Save(user *model.User) (*model.User, error){

	return createTestUser(), nil
}

func (m *UserRepositorySuccessMock) Update(user *model.User) (*model.User, error){

	return createTestUser(), nil
}

func (m *UserRepositorySuccessMock) Delete(user *model.User) (*model.User, error){

	return createTestUser(), nil
}


// エラーを返却するモック
type UserRepositoryErrorMock struct {
}

func (m *UserRepositoryErrorMock) FindByUserId(id uint64) (*model.User, error) {

	return nil, errors.New(createErrorMessage())
}

func (m *UserRepositoryErrorMock) FindByEmail(email string) (*model.User, error) {

	return nil, errors.New(createErrorMessage())
}

func (m *UserRepositoryErrorMock) Save(user *model.User) (*model.User, error){

	return nil, errors.New(createErrorMessage())
}

func (m *UserRepositoryErrorMock) Update(user *model.User) (*model.User, error){

	return nil, errors.New(createErrorMessage())
}

func (m *UserRepositoryErrorMock) Delete(user *model.User) (*model.User, error){

	return nil, errors.New(createErrorMessage())
}

// GetById正常系
func TestGetByIdSuccess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	service := NewUserService(&UserRepositorySuccessMock{})

	user, err := service.GetById(1)

	assert.EqualValues(t, createTestUser(), user, "not Matched User")
	assert.Nil(t, err, "Not Nil Error")
}

// GetById異常系
func TestGetByIdError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	service := NewUserService(&UserRepositoryErrorMock{})

	user, err := service.GetById(1)

	assert.Nil(t, user, "Not Nil User")
	assert.EqualValues(t, err, errors.New(createErrorMessage()), "not Matched Error")
}


// CreateUser正常系
func TestCreateUserSuccess(t *testing.T)  {

	gin.SetMode(gin.TestMode)

	service := NewUserService(&UserRepositorySuccessMock{})

	user, err := service.CreateUser(createTestUser())

	assert.EqualValues(t, createTestUser(), user, "not Matched User")
	assert.Nil(t, err, "Not Nil Error")
}
// CreateUser ユーザー取得でエラー
func TestCreateUserFindUserError(t *testing.T)  {

	gin.SetMode(gin.TestMode)

	service := NewUserService(&UserRepositoryErrorMock{})

	user, err := service.CreateUser(createTestUser())

	assert.Nil(t, user, "not Nil User")
	assert.EqualValues(t, errors.New(createErrorMessage()), err, "Not Matched Error")
}

// UpdateUser正常系
func TestUpdateUserSuccess(t *testing.T)  {

	gin.SetMode(gin.TestMode)

	service := NewUserService(&UserRepositorySuccessMock{})

	user, err := service.UpdateUser(createTestUser())

	assert.EqualValues(t, createTestUser(), user, "not Matched User")
	assert.Nil(t, err, "Not Nil Error")
}

// UpdateUser ユーザー取得でエラー
func TestUpdateUserFindUserError(t *testing.T)  {
	gin.SetMode(gin.TestMode)

	service := NewUserService(&UserRepositoryErrorMock{})

	user, err := service.CreateUser(createTestUser())

	assert.Nil(t, user, "not Nil User")
	assert.EqualValues(t, errors.New(createErrorMessage()), err, "Not Matched Error")
}

// DeleteUser 正常系
func TestDeleteUserSuccess(t *testing.T)  {

	gin.SetMode(gin.TestMode)

	service := NewUserService(&UserRepositorySuccessMock{})

	user, err := service.DeleteUser(1)

	assert.EqualValues(t, createTestUser(), user, "not Matched User")
	assert.Nil(t, err, "Not Nil Error")
}



// テストユーザー作成
func createTestUser() *model.User {

	common := model.CommonModelFields{
		Id:1,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	user := model.User {
		CommonModelFields: common,
		FirstName:"FirstNameTest",
		LastName:"LastNameTest",
		Email:"test@gmail.com",
		Status:model.UserStatusValid,
	}

	return &user
}

// エラーメッセージの返却
func createErrorMessage() string {

	return "Test Error Message"
}