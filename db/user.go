package db

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrNoUser = errors.New("no user")

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FindUserBy(column string, value interface{}) (*User, error) {
	u := &User{}
	wheres := make(map[string]interface{})
	wheres[column] = value

	query := BasicSelectQuery{
		table:   "users",
		columns: []string{"id", "email", "token"},
		wheres:  wheres,
	}

	if err := Select(u, query); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoUser
		}
		return nil, err
	}
	return u, nil
}

func FindUserByEmail(email string) (*User, error) {
	return FindUserBy("email", email)
}

func FindUserById(id int) (*User, error) {
	return FindUserBy("id", id)
}

func FindUserByToken(token string) (*User, error) {
	return FindUserBy("token", token)
}

func CreateUser(email string) (*User, error) {
	token, err := generateUniqueToken()
	if err != nil {
		return nil, err
	}

	colvalues := make(map[string]interface{})
	colvalues["email"] = email
	colvalues["token"] = token

	query := BasicInsertQuery{
		table:     "users",
		colvalues: colvalues,
	}

	_, err = Insert(query)
	if err != nil {
		return nil, err
	}

	return FindUserByEmail(email)
}

func generateUniqueToken() (string, error) {
	var token string
	var count int

	// not sure if this is needed but just to be safe I guess
	for {
		token = randomToken()
		if err := db.Get(&count, "SELECT count(*) FROM users WHERE token=$1", token); err != nil {
			return "", err
		}
		if count == 0 {
			break
		}
	}

	return token, nil
}

func randomToken() string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
