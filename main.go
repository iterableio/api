package main

import (
	"log"
	"net/http"

	_ "github.com/iterableio/api/config"
	"github.com/iterableio/api/core"
	_ "github.com/iterableio/api/db"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", core.InitRouter()))
}
