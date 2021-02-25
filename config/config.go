package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

// ServerConfig struct
type ServerConfig struct {
	AppVersion string
	Port       string
	Mode       string
}

// PostgresConfig struct
type PostgresConfig struct {
	Hostname   string
	Port       string
	Username   string
	Password   string
	DBName     string
	DriverName string
}

// ReadConfig fn
func ReadConfig(filename string) (*viper.Viper, error) {

	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return nil, errors.New("Config file not found")
		}
		// Config file was found but another error was produced
		return nil, err
	}

	return v, nil

}

// GetConfig get config by pathname
func GetConfig(pathname string) (*Config, error) {

	configFile, err := ReadConfig(pathname)

	if err != nil {
		return nil, err
	}

	var config Config

	if err = configFile.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil

}

// GetConfigByName fn
func GetConfigByName(name string) string {
	pathname := "config-dev"

	if name == "production" {
		pathname = "config-prod"
	}

	return fmt.Sprintf("./config/%s", pathname)
}
