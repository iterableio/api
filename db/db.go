package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/iterableio/api/config"
)

var db *sqlx.DB

func ConnectSQL() error {
	var err error
	db, err = sqlx.Connect("postgres", buildDataSourceName())
	return err
}

func buildDataSourceName() string {
	return fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
		config.Global.Postgres.User,
		config.Global.Postgres.Password,
		config.Global.Postgres.DBName,
		config.Global.Postgres.SSLMode)
}

func init() {
	log.Println("Starting DB")
	if err := ConnectSQL(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("DB started")
}
