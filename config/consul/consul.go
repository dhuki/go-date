package consul

import (
	"context"
	"os"

	"github.com/dhuki/go-date/config"
	"github.com/dhuki/go-date/pkg/logger"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// InitialiseRemote sets the remote configuration
func InitialiseRemote(ctx context.Context, v *viper.Viper) error {
	var consulEndpoint string
	if consulEnv := os.Getenv("CONSUL_URL"); consulEnv != "" {
		consulEndpoint = consulEnv
	}
	logger.Info(ctx, "Initials remote config", "Initialising remote config, consul endpoint: %s", consulEndpoint)
	err := v.AddRemoteProvider("consul", consulEndpoint, "GO_DATE")
	if err != nil {
		logger.Warn(ctx, "initials remote config", "error add remote config provider, error: %v", err)
		return err
	}
	v.SetConfigType("yaml")
	return v.ReadRemoteConfig()
}

// InitialiseFileAndEnv sets the file configuration and the environment
func InitialiseFileAndEnv(ctx context.Context, v *viper.Viper, configName string) error {
	v.SetConfigName(configName)
	for _, path := range searchPath {
		v.AddConfigPath(path)
	}
	v.AutomaticEnv()
	return v.ReadInConfig()
}

// LoadConfig load the configuration from remote or local
func InitConsul(ctx context.Context, env string) {
	v := viper.New()
	if err := InitialiseRemote(ctx, v); err != nil {
		logger.Warn(ctx, "initials remote config", "No remote server configured will load configuration from file and environment variables, %v", err)
		logger.Info(ctx, "initials local config", "Looking for local config file '%s' in search paths", configName[env])
		if err := InitialiseFileAndEnv(ctx, v, configName[env]); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				logger.Fatal(ctx, "initials local config", "No '%s' file found on search paths. Will either use environment variables or defaults", configName[env])
			} else {
				logger.Fatal(ctx, "initials local config", "error occured during loading config, error: %v", err)
			}
		}
	}
	err := v.Unmarshal(&config.Conf)
	if err != nil {
		logger.Fatal(ctx, "marshalling config", "Error occured during unmarshalling config: %s", err.Error())
	}
}
