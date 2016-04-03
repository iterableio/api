package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndFindUser(t *testing.T) {
	assert := assert.New(t)

	user, err := CreateUser("swag@swag.swag")
	assert.Nil(err)

	userAgain, err := FindUserById(user.Id)
	assert.Nil(err)

	assert.Equal(user.Id, userAgain.Id)
	assert.Equal(user.Email, userAgain.Email)
	assert.Equal(user.Token, userAgain.Token)
}

func TestFailCreateWithSameEmail(t *testing.T) {
	assert := assert.New(t)

	email := "same@email.com"

	_, err := CreateUser(email)
	assert.Nil(err)

	_, err = CreateUser(email)
	assert.NotNil(err)
}
