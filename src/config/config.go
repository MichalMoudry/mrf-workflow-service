package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port             int
	ConnectionString string
	RunWithDapr      bool
	Environment
}

// This function reads app's configuration from a config file.
func ReadCfgFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	connectionString := os.Getenv("DB_CONN")
	if connectionString == "" {
		connectionString = viper.GetString("connection_string")
	}

	return Config{
		Port:             viper.GetInt("port"),
		ConnectionString: connectionString,
		Environment:      Environment(viper.GetString("environment")),
	}, nil
}
