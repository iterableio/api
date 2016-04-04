package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iterableio/api/common"
	"github.com/iterableio/api/db"
)

var server *httptest.Server

func init() {
	server = httptest.NewServer(InitRouter())
}

func getUrl(route string) string {
	return fmt.Sprintf("%s/api/v1%s", server.URL, route)
}

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	resp, err := http.Get(getUrl("/"))
	assert.Nil(err)
	assert.Equal(resp.StatusCode, 200)

	actual, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)
	assert.Equal("le iterable api", string(actual))
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	email := common.RandomEmail()
	payload := []byte(fmt.Sprintf(`{"email": "%s"}`, email))

	resp, err := http.Post(getUrl("/users"), "application/json", bytes.NewBuffer(payload))
	assert.Nil(err)
	assert.Equal(resp.StatusCode, 200)

	var user db.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	assert.Nil(err)
	assert.Equal(user.Email, email)
}

func TestCreateUserSameEmail(t *testing.T) {
	assert := assert.New(t)

	statusCodes := []int{200, 400}
	email := common.RandomEmail()
	payload := []byte(fmt.Sprintf(`{"email": "%s"}`, email))

	for _, code := range statusCodes {
		resp, err := http.Post(getUrl("/users"), "application/json", bytes.NewBuffer(payload))
		assert.Nil(err)
		assert.Equal(resp.StatusCode, code)
	}
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	user, err := db.CreateUser(common.RandomEmail())
	assert.Nil(err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", getUrl(fmt.Sprintf("/users/%d", user.Id)), nil)
	assert.Nil(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", user.Token))

	resp, err := client.Do(req)
	assert.Equal(resp.StatusCode, 200)

	var userAgain db.User
	err = json.NewDecoder(resp.Body).Decode(&userAgain)

	assert.Nil(err)
	assert.Equal(user.Id, userAgain.Id)
	assert.Equal(user.Email, userAgain.Email)
	assert.Equal(user.Token, userAgain.Token)
}

func TestFailGetUserNoAuth(t *testing.T) {
	assert := assert.New(t)

	user, err := db.CreateUser(common.RandomEmail())
	assert.Nil(err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", getUrl(fmt.Sprintf("/users/%d", user.Id)), nil)
	assert.Nil(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	assert.Equal(resp.StatusCode, 401)
}

func TestGetWrongUser(t *testing.T) {
	assert := assert.New(t)

	user, err := db.CreateUser(common.RandomEmail())
	assert.Nil(err)

	target, err := db.CreateUser(common.RandomEmail())
	assert.Nil(err)

	client := &http.Client{}
	req, err := http.NewRequest("GET", getUrl(fmt.Sprintf("/users/%d", target.Id)), nil)
	assert.Nil(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", user.Token))

	resp, err := client.Do(req)
	assert.Equal(resp.StatusCode, 401)
}
