package main

import (
	"fmt"

	_ "github.com/iterableio/api/config"
	"github.com/iterableio/api/db"
)

func main() {
	fmt.Printf("Starting db\n")
	db.ConnectSQL()
	fmt.Printf("db started\n")
}
