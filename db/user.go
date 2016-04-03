package db

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id    int
	Email string
	Token string
}

func FindUserBy(column string, value interface{}) (User, error) {
	var u User
	err := db.QueryRowx(fmt.Sprintf("SELECT id, email, token FROM users WHERE %v=$1", column), value).StructScan(&u)
	return u, err
}

func FindUserByEmail(email string) (User, error) {
	return FindUserBy("email", email)
}

func FindUserById(id int) (User, error) {
	return FindUserBy("id", id)
}

func CreateUser(email string) (User, error) {
	var u User

	token, err := generateUniqueToken()
	if err != nil {
		return u, err
	}

	_, err = db.Exec(`INSERT INTO users (email, token) VALUES ($1, $2)`, email, token)
	if err != nil {
		return u, err
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
