package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port             int
	ConnectionString string
	RunWithDapr      bool
	RunWithFirebase  bool
	Environment
}

// This function reads app's configuration from a config file.
func ReadCfgFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	connectionString := os.Getenv("PROD_DB_STRING")
	if connectionString == "" {
		connectionString = viper.GetString("connection_string")
	}

	return Config{
		Port:             viper.GetInt("port"),
		ConnectionString: connectionString,
		RunWithDapr:      viper.GetBool("run_with_dapr"),
		RunWithFirebase:  viper.GetBool("run_with_firebase"),
		Environment:      Environment(viper.GetString("environment")),
	}, nil
}
