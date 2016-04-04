package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/iterableio/api/db"
)

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// this is actually horrible
	targetId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		WriteErrorBadRequest(w, err)
		return
	}
	if currentUser := getCurrentUser(r); currentUser.Id != targetId {
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
		WriteErrorBadRequest(w, err)
		return
	}
	WriteResponse(w, user)
}
