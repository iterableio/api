package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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
			// do things
			WriteErrorUnauthorized(w, errors.New("Missing/incorrect auth header"))
			return
		}
		user, err := db.FindUserByToken(payload[1])
		if err != nil {
			WriteErrorUnauthorized(w, errors.New("User doesn't exist"))
			return
		}
		context.Set(r, "user", user)
		h(w, r, ps)
		context.Clear(r)
	}
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	targetId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		WriteErrorBadRequest(w, err)
		return
	}

	if currentUser := context.Get(r, "user").(db.User); currentUser.Id != targetId {
		WriteErrorUnauthorized(w, errors.New("You dont have permission to view this"))
		return
	}

	target, err := db.FindUserById(targetId)
	if err != nil {
		WriteErrorBadRequest(w, err)
		return
	}
	WriteResponse(w, target)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req struct {
		Email string
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorInternal(w, err)
		return
	}
	if !ValidEmail(req.Email) {
		WriteErrorBadRequest(w, errors.New("Invalid email"))
		return
	}
	user, err := db.CreateUser(req.Email)
	if err != nil {
		WriteErrorInternal(w, err)
		return
	}
	WriteResponse(w, user)
}
