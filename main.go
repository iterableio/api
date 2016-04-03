package main

import (
	"log"
	"net/http"

	"github.com/iterableio/api/core"

	_ "github.com/iterableio/api/config"
	"github.com/iterableio/api/db"
)

func main() {
	log.Println("Starting DB")
	if err := db.ConnectSQL(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("DB started")
	log.Fatal(http.ListenAndServe(":8080", core.InitRouter()))
}
