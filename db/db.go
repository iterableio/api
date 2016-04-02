package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/iterableio/api/config"
)

var DB *sqlx.DB

func ConnectSQL() error {
	var err error
	DB, err = sqlx.Connect("postgres", buildDataSourceName())
	return err
}

func buildDataSourceName() string {
	return fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
		config.Global.Postgres.User,
		config.Global.Postgres.Password,
		config.Global.Postgres.DBName,
		config.Global.Postgres.SSLMode)
}
