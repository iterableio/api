package core

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"

	"github.com/iterableio/api/db"
)

func auth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")

		// should have format `token {TOKEN}`
		payload := strings.Split(authHeader, " ")
		if len(payload) != 2 || payload[0] != "token" {
			WriteErrorUnauthorized(w, errors.New("Missing/incorrect auth header"))
			return
		}
		user, err := db.FindUserByToken(payload[1])
		if err != nil {
			WriteErrorBadRequest(w, err)
			return
		}
		setCurrentUser(r, user)
		h(w, r, ps)
		context.Clear(r)
	}
}

func getCurrentUser(r *http.Request) *db.User {
	return context.Get(r, "user").(*db.User)
}

func setCurrentUser(r *http.Request, u *db.User) {
	context.Set(r, "user", u)
}
