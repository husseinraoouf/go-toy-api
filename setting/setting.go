package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Type             string
	ConnectionString string
	Path             string
}

var Config config

func LoadConfig() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // add current directory to config search paths

	var err error
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	return nil
}
