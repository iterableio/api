package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDataSourceName(t *testing.T) {
	assert := assert.New(t)

	dataSourceName := buildDataSourceName()

	assert.Equal(dataSourceName,
		"user=postgres password=postgres dbname=iterable_test sslmode=disable",
		"Should be consistent with test.yaml")
}

func TestConnectSQL(t *testing.T) {
	assert := assert.New(t)

	ConnectSQL()

	assert.NotNil(DB, "Should not be nil, should be set to test DB")
}
