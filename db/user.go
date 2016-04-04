package db

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
)

var ErrNoUser = errors.New("no user")

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FindUserBy(column string, value interface{}) (*User, error) {
	u := &User{}
	q := fmt.Sprintf("SELECT id, email, token FROM users WHERE %v=$1", column)

	if err := db.Get(u, q, value); err != nil {
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

	_, err = db.Exec(`INSERT INTO users (email, token) VALUES ($1, $2)`, email, token)
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
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
