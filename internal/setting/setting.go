package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

// config represents the app config.
type config struct {
	Database databaseConfig
}

// databaseConfig represents the database config.
type databaseConfig struct {
	Type             string
	ConnectionString string
	Path             string
}

var Config config

// LoadConfig loads config from files.
func LoadConfig() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // add current directory to config search paths

	var err error

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %w", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %w", err)
	}

	return nil
}
