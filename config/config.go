package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port             int
	ConnectionString string
}

// This function reads app's configuration from a config file.
func ReadCfgFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		connectionString = viper.GetString("ConnectionString")
	}

	return Config{
		Port:             viper.GetInt("Port"),
		ConnectionString: connectionString,
	}, nil
}
