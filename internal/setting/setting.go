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

// LoadConfig loads default config files.
func LoadConfig() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // add current directory to config search paths

	if err := loadConfig(); err != nil {
		return err
	}

	return nil
}

// LoadTestConfigAtPath loads test config files at given path.
func LoadTestConfigAtPath(path string) error {
	viper.SetConfigName("config.test") // name of config file (without extension)
	viper.AddConfigPath(path)          // add current directory to config search paths

	if err := loadConfig(); err != nil {
		return err
	}

	return nil
}

// LoadConfig loads config files.
func loadConfig() error {
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
