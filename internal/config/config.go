package config

import (
	"flag"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	GRPC      GRPCConfig
	BFALimits BFALimitsConfig
}

type GRPCConfig struct {
	Port int
}

type BFALimitsConfig struct {
	Login          int
	Password       int
	IP             int
	Period         int
	BucketLifetime int
}

func Load() (*Config, error) {
	var cfg Config

	data, err := os.ReadFile(fetchConfigPath())

	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func fetchConfigPath() string {
	res := os.Getenv("CONFIG_PATH")

	if res == "" {
		flag.StringVar(&res, "config", defaultConfigPath(), "path to config file")
		flag.Parse()
	}

	return res
}

func defaultConfigPath() string {
	_, filename, _, _ := runtime.Caller(0)

	return filepath.Join(filename, "../../../config/development.toml")
}
