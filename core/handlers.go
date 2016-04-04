package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/iterableio/api/db"
	"github.com/julienschmidt/httprouter"
)

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		WriteErrorBadRequest(w, err)
		return
	}
	user, err := db.FindUserById(userId)
	if err != nil {
		WriteErrorBadRequest(w, err)
		return
	}
	WriteResponse(w, user)
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
