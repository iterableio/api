package core

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/", index)
	router.GET("/api/users/:userId", auth(getUser))
	router.POST("/api/users", createUser)

	return router
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "le iterable api")
}
