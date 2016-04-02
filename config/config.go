package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Environment variable keys
const (
	environmentKey = "ITERABLE_ENVIRONMENT"
	configDirKey   = "ITERABLE_CONFIG_DIR"
)

// Other constants
const (
	configDir = "config"
)

var Global IterableConfig

type SQLConfig struct {
	User     string `yaml:user`
	Password string `yaml:password`
	DBName   string `yaml:dbname`
	SSLMode  string `yaml:sslmode`
}

type IterableConfig struct {
	Postgres SQLConfig `yaml:postgres`
}

type ConfigParser func(in []byte, out interface{}) error

func init() {
	configPath := getConfigPath()
	if err := load(&Global, configPath, yaml.Unmarshal); err != nil {
		log.Fatal(err)
	}
}

// Read and Unmarshall config file
func load(conf *IterableConfig, path string, configParserFunc ConfigParser) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := configParserFunc(data, conf); err != nil {
		return err
	}
	return nil
}

// Get config file path from environment variables
func getConfigPath() string {
	realConfigDir := configDir
	if envConfigDir := os.Getenv(configDirKey); envConfigDir != "" {
		realConfigDir = envConfigDir
	}
	fileName := fmt.Sprintf("%s.yaml", os.Getenv(environmentKey))
	return filepath.Join(realConfigDir, fileName)
}
