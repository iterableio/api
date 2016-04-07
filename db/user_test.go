package db

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iterableio/api/common"
)

func TestCreateAndFindUser(t *testing.T) {
	assert := assert.New(t)

	user, err := CreateUser(common.RandomEmail())
	assert.Nil(err)

	userAgain, err := FindUserById(user.Id)
	assert.Nil(err)

	assert.Equal(user.Id, userAgain.Id)
	assert.Equal(user.Email, userAgain.Email)
	assert.Equal(user.Token, userAgain.Token)
}

func TestFailCreateWithSameEmail(t *testing.T) {
	assert := assert.New(t)

	email := common.RandomEmail()

	_, err := CreateUser(email)
	assert.Nil(err)

	_, err = CreateUser(email)
	assert.NotNil(err)
}
