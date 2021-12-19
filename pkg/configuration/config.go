package configuration

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

type ApplicationConfig struct {
	Web struct {
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
		IdleTimeout     time.Duration
		ShutdownTimeout time.Duration
		APIHost         string
		DebugHost       string
	}
}

type Config interface {
	GetConfig() (*ApplicationConfig, error)
}

type config struct{}

func (config *config) GetConfig(path string) (*ApplicationConfig, error) {
	configuration := ApplicationConfig{}
	env := getGoEnv()

	viperInstance := getViperInstance(path)
	err := viperInstance.ReadInConfig()

	if err != nil {
		return nil, err
	}

	sub := viperInstance.Sub(env)
	err = sub.Unmarshal(&configuration)

	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func CreateConfigInstance() *config {
	return &config{}
}

func getViperInstance(path string) *viper.Viper {
	viperInstance := viper.New()
	viperInstance.SetConfigFile(path)
	return viperInstance
}

func getGoEnv() string {
	env := os.Getenv("GO_ENV")
	if env != "" {
		return env
	}
	return "stage"
}
