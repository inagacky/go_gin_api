package controller


import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/inagacky/go_gin_api/src/api/domain/model"
	"github.com/inagacky/go_gin_api/src/api/usecase"
	usecaseUser "github.com/inagacky/go_gin_api/src/api/usecase/user"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


var timeNow = time.Now()

// 正常値を返却するモック
type UserServiceSuccessMock struct {
}

func (m *UserServiceSuccessMock) GetById(id uint64) (*model.User, error) {

	user := createTestUser()
	return user, nil

}

func (m *UserServiceSuccessMock) CreateUser(user *model.User) (*model.User, error) {

	testUser := createTestUser()
	return testUser, nil
}

func (m *UserServiceSuccessMock) UpdateUser(user *model.User) (*model.User, error){

	testUser := createTestUser()
	return testUser, nil
}

func (m *UserServiceSuccessMock) DeleteUser(id uint64) (*model.User, error){
	testUser := createTestUser()
	return testUser, nil
}

// エラーを返却するモック
type UserServiceErrorMock struct {
}

func (m *UserServiceErrorMock) GetById(id uint64) (*model.User, error) {

	err := errors.New(createErrorMessage())
	return nil, err

}

func (m *UserServiceErrorMock) CreateUser(user *model.User) (*model.User, error) {

	err := errors.New(createErrorMessage())
	return nil, err
}

func (m *UserServiceErrorMock) UpdateUser(user *model.User) (*model.User, error){

	err := errors.New(createErrorMessage())
	return nil, err
}

func (m *UserServiceErrorMock) DeleteUser(id uint64) (*model.User, error){

	err := errors.New(createErrorMessage())
	return nil, err
}

type testPattern struct{
	Request testRequest
	Response testResponse
}

// Requestのユーザー情報
type testRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// レスポンス期待値
type testResponse struct {
	Status int
	Message string
}


// GetUser正常系
func TestGetUserSuccess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	userCo := NewUserController(&UserServiceSuccessMock{})
	r.GET("/:id", userCo.GetUser)
	req, _ := http.NewRequest("GET", "/1", nil)
	r.ServeHTTP(w, req)

	response := testResponse{
		Status: http.StatusOK,
		Message: "",
	}

	// ステータスコードのチェック
	assert.Equal(t, response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

	// レスポンス結果のチェック
	commonResponse := usecase.CommonResponse{}
	jVal, _ := json.Marshal(commonResponse.CreateSuccessResponse(usecaseUser.GetUserResponse{User:createTestUser()}))
	assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
}

// GetUserリクエスト情報に誤りがある場合
func TestGetUserRequestError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	userCo := NewUserController(&UserServiceSuccessMock{})
	r.GET("/:id", userCo.GetUser)
	req, _ := http.NewRequest("GET", "/-1", nil)
	r.ServeHTTP(w, req)

	response := testResponse{
		Status: http.StatusBadRequest,
		Message: "Key: 'GetUserRequest.Id' Error:Field validation for 'Id' failed on the 'number' tag",
	}

	// ステータスコードのチェック
	assert.Equal(t, response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

	// レスポンス結果のチェック
	err := errors.New(response.Message)
	commonResponse := usecase.CommonResponse{}
	jVal, _ := json.Marshal(commonResponse.CreateValidateErrorResponse(err.Error()))
	assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
}

// GetUser Service層でのエラー
func TestGetUserServiceError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	userCo := NewUserController(&UserServiceErrorMock{})
	r.GET("/:id", userCo.GetUser)
	req, _ := http.NewRequest("GET", "/1", nil)
	r.ServeHTTP(w, req)

	response := testResponse{
		Status: http.StatusBadRequest,
		Message: errors.New(createErrorMessage()).Error(),
	}

	// ステータスコードのチェック
	assert.Equal(t, response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

	// レスポンス結果のチェック
	commonResponse := usecase.CommonResponse{}
	jVal, _ := json.Marshal(commonResponse.CreateSQLErrorResponse(response.Message))
	assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
}


// CreateUser正常系
func TestCreateUserSuccess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	// リクエスト情報
	requestUser := testRequest{
		FirstName: "TestFirstName",
		LastName: "TestLastName",
		Email:"test@sample.co.jp",
	}
	// レスポンス情報
	response := testResponse{
		Status: http.StatusOK,
		Message: "",
	}

	postVal, _ := json.Marshal(requestUser)

	userCo := NewUserController(&UserServiceSuccessMock{})
	r.POST("/", userCo.CreateUser)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(postVal))
	r.ServeHTTP(w, req)

	// ステータスコードのチェック
	assert.Equal(t, response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

	// レスポンス結果のチェック
	commonResponse := usecase.CommonResponse{}
	jVal, _ := json.Marshal(commonResponse.CreateSuccessResponse(usecaseUser.CreateUserResponse{User:createTestUser()}))
	assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
}

// CreateUserリクエスト情報に誤りがある場合
func TestCreateUserRequestError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	tests := []testPattern{
		{
			// FirstNameがブランク
			testRequest{ FirstName: "", LastName: "TestLastName", Email:"test@sample.co.jp" },
			testResponse{Status: http.StatusBadRequest, Message: "Key: 'CreateUserRequest.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag"},
		},
		{
			// LastNameがブランク
			testRequest{ FirstName: "TestFirstName", LastName: "", Email:"test@sample.co.jp" },
			testResponse{Status: http.StatusBadRequest, Message: "Key: 'CreateUserRequest.LastName' Error:Field validation for 'LastName' failed on the 'required' tag"},
		},
		{
			// Emailがブランク
			testRequest{ FirstName: "TestFirstName", LastName: "TestLastName", Email:"" },
			testResponse{Status: http.StatusBadRequest, Message: "Key: 'CreateUserRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"},
		},
		{
			// Emailのフォーマット不正
			testRequest{ FirstName: "TestFirstName", LastName: "TestLastName", Email:"test" },   // Emailのフォーマット不正
			testResponse{Status: http.StatusBadRequest, Message: "Key: 'CreateUserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"},
		},

	}

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	userCo := NewUserController(&UserServiceSuccessMock{})
	r.POST("/", userCo.CreateUser)

	for _, test := range tests {

		w := httptest.NewRecorder()
		postVal, _ := json.Marshal(test.Request)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(postVal))
		r.ServeHTTP(w, req)

		assert.Equal(t, test.Response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

		commonResponse := usecase.CommonResponse{}
		jVal, _ := json.Marshal(commonResponse.CreateValidateErrorResponse(test.Response.Message))

		assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
	}
}


// CreateUser Service層でのエラー
func TestCreateUserServiceError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	tests := []testPattern{
		{
			// FirstNameがnil
			testRequest{ FirstName: "TestFirstName", LastName: "TestLastName", Email:"test@sample.co.jp" },
			testResponse{Status: http.StatusBadRequest, Message: errors.New(createErrorMessage()).Error() },
		},
	}

	userCo := NewUserController(&UserServiceErrorMock{})
	r.POST("/", userCo.CreateUser)

	for _, test := range tests {

		postVal, _ := json.Marshal(test.Request)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(postVal))
		r.ServeHTTP(w, req)

		assert.Equal(t, test.Response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

		commonResponse := usecase.CommonResponse{}
		jVal, _ := json.Marshal(commonResponse.CreateSQLErrorResponse(test.Response.Message))
		assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
	}
}



// UpateUser正常系
func TestUpdateUserSuccess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	// リクエスト情報
	requestUser := testRequest{
		FirstName: "TestFirstName",
		LastName: "TestLastName",
		Email:"test@sample.co.jp",
	}
	// レスポンス情報
	response := testResponse{
		Status: http.StatusOK,
		Message: "",
	}

	postVal, _ := json.Marshal(requestUser)

	userCo := NewUserController(&UserServiceSuccessMock{})
	r.PUT("/:id", userCo.UpdateUser)
	req, _ := http.NewRequest("PUT", "/1", bytes.NewBuffer(postVal))
	r.ServeHTTP(w, req)

	// ステータスコードのチェック
	assert.Equal(t, response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

	// レスポンス結果のチェック
	commonResponse := usecase.CommonResponse{}
	jVal, _ := json.Marshal(commonResponse.CreateSuccessResponse(usecaseUser.UpdateUserResponse{User:createTestUser()}))
	assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
}

// UpdateUserリクエスト情報に誤りがある場合
func TestUpdateUserRequestError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	tests := []testPattern{
		{
			// FirstNameがブランク
			testRequest{ FirstName: "", LastName: "TestLastName", Email:"test@sample.co.jp" },
			testResponse{Status: http.StatusBadRequest, Message: "Key: 'UpdateUserRequest.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag"},
		},
		{
			// LastNameがブランク
			testRequest{ FirstName: "TestFirstName", LastName: "", Email:"test@sample.co.jp" },
			testResponse{Status: http.StatusBadRequest, Message: "Key: 'UpdateUserRequest.LastName' Error:Field validation for 'LastName' failed on the 'required' tag"},
		},
		{
			// Emailがブランク
			testRequest{ FirstName: "TestFirstName", LastName: "TestLastName", Email:"" },
			testResponse{Status: http.StatusBadRequest, Message: "Key: 'UpdateUserRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"},
		},
		{
			// Emailのフォーマット不正
			testRequest{ FirstName: "TestFirstName", LastName: "TestLastName", Email:"test" },   // Emailのフォーマット不正
			testResponse{Status: http.StatusBadRequest, Message: "Key: 'UpdateUserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"},
		},

	}

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	userCo := NewUserController(&UserServiceSuccessMock{})
	r.PUT("/:id", userCo.UpdateUser)

	for _, test := range tests {

		w := httptest.NewRecorder()
		postVal, _ := json.Marshal(test.Request)
		req, _ := http.NewRequest("PUT", "/1", bytes.NewBuffer(postVal))
		r.ServeHTTP(w, req)

		assert.Equal(t, test.Response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

		commonResponse := usecase.CommonResponse{}
		jVal, _ := json.Marshal(commonResponse.CreateValidateErrorResponse(test.Response.Message))

		assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
	}
}


// UpdateUser Service層でのエラー
func TestUpdateUserServiceError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	tests := []testPattern{
		{
			// FirstNameがnil
			testRequest{ FirstName: "TestFirstName", LastName: "TestLastName", Email:"test@sample.co.jp" },
			testResponse{Status: http.StatusBadRequest, Message: errors.New(createErrorMessage()).Error() },
		},
	}

	userCo := NewUserController(&UserServiceErrorMock{})
	r.POST("/", userCo.CreateUser)

	for _, test := range tests {

		postVal, _ := json.Marshal(test.Request)
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(postVal))
		r.ServeHTTP(w, req)

		assert.Equal(t, test.Response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

		commonResponse := usecase.CommonResponse{}
		jVal, _ := json.Marshal(commonResponse.CreateSQLErrorResponse(test.Response.Message))
		assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
	}
}


// DeleteUser 正常系
func TestDeleteUserSuccess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	userCo := NewUserController(&UserServiceSuccessMock{})
	r.DELETE("/:id", userCo.DeleteUser)
	req, _ := http.NewRequest("DELETE", "/1", nil)
	r.ServeHTTP(w, req)

	response := testResponse{
		Status: http.StatusOK,
		Message: "",
	}

	// ステータスコードのチェック
	assert.Equal(t, response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

	// レスポンス結果のチェック
	commonResponse := usecase.CommonResponse{}
	jVal, _ := json.Marshal(commonResponse.CreateSuccessResponse(usecaseUser.DeleteUserResponse{User:createTestUser()}))
	assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
}

// DeleteUserリクエスト情報に誤りがある場合
func TestDeleteUserRequestError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	userCo := NewUserController(&UserServiceSuccessMock{})
	r.DELETE("/:id", userCo.DeleteUser)
	req, _ := http.NewRequest("DELETE", "/-1", nil)
	r.ServeHTTP(w, req)

	response := testResponse{
		Status: http.StatusBadRequest,
		Message: "Key: 'DeleteUserRequest.Id' Error:Field validation for 'Id' failed on the 'number' tag",
	}

	// ステータスコードのチェック
	assert.Equal(t, response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

	// レスポンス結果のチェック
	err := errors.New(response.Message)
	commonResponse := usecase.CommonResponse{}
	jVal, _ := json.Marshal(commonResponse.CreateValidateErrorResponse(err.Error()))
	assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
}

// DeleteUser Service層でのエラー
func TestDeleteUserServiceError(t *testing.T) {

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	userCo := NewUserController(&UserServiceErrorMock{})
	r.DELETE("/:id", userCo.DeleteUser)
	req, _ := http.NewRequest("DELETE", "/1", nil)
	r.ServeHTTP(w, req)

	response := testResponse{
		Status: http.StatusBadRequest,
		Message: errors.New(createErrorMessage()).Error(),
	}

	// ステータスコードのチェック
	assert.Equal(t, response.Status, w.Code, "not Matched HttpCode: %v", w.Code)

	// レスポンス結果のチェック
	commonResponse := usecase.CommonResponse{}
	jVal, _ := json.Marshal(commonResponse.CreateSQLErrorResponse(response.Message))
	assert.Equal(t, w.Body.String(), string(jVal), "not Matched ResponseBody: %v", w.Body.String())
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