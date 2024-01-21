package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Model struct {
	App    string
	AppVer string
	Env    string
	Http   HttpConfig
}

type HttpConfig struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

func (m *Model) LoadConfig(path string) {
	viper.AddConfigPath(filepath.Dir(path))
	viper.SetConfigName(filepath.Base(path))
	viper.SetConfigType(filepath.Ext(path)[1:])

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	err = viper.Unmarshal(m)
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
}

func getEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}
	return value
}
