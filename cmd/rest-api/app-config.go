package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	LogLevel                  string `mapstructure:"LOG_LEVEL"`
	Port                      uint16 `mapstructure:"PORT"`
	MpdHost                   string `mapstructure:"MPD_HOST"`
	MpdPort                   uint16 `mapstructure:"MPD_PORT"`
	MpdPassword               string `mapstructure:"MPD_PASSWORD"`
	MaxBatchCommand           uint16 `mapstructure:"MAX_BATCH_COMMAND"`
	CommandReadIntervalMillis uint16 `mapstructure:"COMMAND_READ_INTERVAL_MILLIS"`
	MpdPoolSize               uint8  `mapstructure:"MPD_POOL_SIZE"`
	PingIntervalSeconds       uint8  `mapstructure:"PING_INTERVAL_SECONDS"`
}

func loadConfig() (*AppConfig, error) {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, err
		}
	}
	result := &AppConfig{}
	bindEnvs(result)
	err := viper.Unmarshal(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func bindEnvs(config interface{}) {
	t := reflect.TypeOf(config)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("mapstructure")
		if tag == "" {
			continue
		}
		if err := viper.BindEnv(tag); err != nil {
			fmt.Printf("failed to bind env for %s: %v\n", tag, err)
		}
	}
}
