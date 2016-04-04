package core

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/", index)
	router.GET("/api/v1/users/:userId", auth(getUser))
	router.POST("/api/v1/users", createUser)

	return router
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "le iterable api")
}
