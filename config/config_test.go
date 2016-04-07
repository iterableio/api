package config

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func mockUnmarshal(in []byte, out interface{}) error {
	return errors.New("")
}

func TestLoad(t *testing.T) {
	assert := assert.New(t)

	conf := IterableConfig{}
	err := load(&conf, "test.yaml", yaml.Unmarshal)
	assert.Nil(err, "There was no error")

	same := "This should be the same as test.yaml"
	assert.Equal(conf.Postgres.DBName, "iterable_test", same)
	assert.Equal(conf.Postgres.SSLMode, "disable", same)
}

func TestLoadFileMissing(t *testing.T) {
	assert := assert.New(t)

	conf := IterableConfig{}
	err := load(&conf, "t2st.yaml", yaml.Unmarshal)

	assert.Error(err, "This file should not exist")
}

func TestLoadParseFail(t *testing.T) {
	assert := assert.New(t)

	conf := IterableConfig{}
	err := load(&conf, "test.yaml", mockUnmarshal)

	assert.Error(err, "This should return an error")
}

func TestGetConfigPathFromEnv(t *testing.T) {
	assert := assert.New(t)
	oldConfigDir := os.Getenv(configDirKey)
	defer os.Setenv(configDirKey, oldConfigDir)
	os.Setenv(configDirKey, "/hello")

	path := getConfigPath()

	assert.Equal(path, "/hello/test.yaml")
}

func TestGetConfigPathDefault(t *testing.T) {
	assert := assert.New(t)
	oldConfigDir := os.Getenv(configDirKey)
	defer os.Setenv(configDirKey, oldConfigDir)
	os.Setenv(configDirKey, "")

	path := getConfigPath()

	assert.Equal(path, "config/test.yaml")
}
