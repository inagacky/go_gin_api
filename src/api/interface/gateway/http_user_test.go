package gateway

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestConvertUserModel(t *testing.T) {

	// CreateUserリクエストのテスト
	createUserRequest := CreateUserRequest{
		FirstName:"TestFirstName",
		LastName:"TestLastName",
		Email:"test@gmail.com",
	}

	user := createUserRequest.ConvertUserModel()

	assert.Equal(t, createUserRequest.FirstName, user.FirstName, "not Matched FirstName")
	assert.Equal(t, createUserRequest.LastName, user.LastName, "not Matched LastName")
	assert.Equal(t, createUserRequest.Email, user.Email, "not Matched Email")

	// UpdateUserリクエストのテスト
	updateUserRequest := UpdateUserRequest{
		Id: "1",
		FirstName:"TestFirstName",
		LastName:"TestLastName",
		Email:"test@gmail.com",
	}
	user = updateUserRequest.ConvertUserModel()

	id, _ := strconv.ParseUint(updateUserRequest.Id, 10 ,64)
	assert.Equal(t, id, user.Id, "not Matched Id")
	assert.Equal(t, updateUserRequest.FirstName, user.FirstName, "not Matched FirstName")
	assert.Equal(t, updateUserRequest.LastName, user.LastName, "not Matched LastName")
	assert.Equal(t, updateUserRequest.Email, user.Email, "not Matched Email")
}
