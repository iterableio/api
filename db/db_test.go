package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iterableio/api/config"
)

func TestBuildDataSourceName(t *testing.T) {
	assert := assert.New(t)

	dataSourceName := buildDataSourceName()

	assert.Equal(dataSourceName,
		fmt.Sprintf("user=%v password=%v dbname=%v sslmode=%v",
			config.Global.Postgres.User,
			config.Global.Postgres.Password,
			config.Global.Postgres.DBName,
			config.Global.Postgres.SSLMode),
		"Should be consistent with test.yaml")
}

func TestConnectSQL(t *testing.T) {
	assert := assert.New(t)

	ConnectSQL()

	assert.NotNil(DB, "Should not be nil, should be set to test DB")
}
