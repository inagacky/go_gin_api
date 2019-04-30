package model


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueCopy(t *testing.T) {

	orgUser := User{
		FirstName: "TestFirstName",
		LastName: "TestLastName",
		Email: "test@gmail.com",
	}
	copyUser := User{}
	copyUser.ValueCopy(&orgUser)

	assert.Equal(t, copyUser.FirstName, orgUser.FirstName, "not Matched User Model FirstName")
	assert.Equal(t, copyUser.LastName, orgUser.LastName, "not Matched User Model LastName")
	assert.Equal(t, copyUser.Email, orgUser.Email, "not Matched User Model Email")

}