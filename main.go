package main

import (
	"log"
	"net/http"

	"github.com/iterableio/api/core"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", core.InitRouter()))
}
