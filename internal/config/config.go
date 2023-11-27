package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Token        string `envconfig:"TOKEN"`
	DatabasePath string `envconfig:"DATABASE_PATH"`
}

func NewAppConfig() (*AppConfig, error) {
	appConfig := &AppConfig{}

	err := envconfig.Process("", appConfig)
	if err != nil {
		return nil, fmt.Errorf("error on parse env variables: %v", err)
	}

	return appConfig, nil
}

func (c *AppConfig) GetToken() string {
	return c.Token
}

func (c *AppConfig) GetDatabasePath() string {
	return c.DatabasePath
}
