package config

import (
	"github.com/BurntSushi/toml"
	"github.com/osamikoyo/pitition/pkg/loger"
)

type Config struct {
	Hostname     string `toml:"hostname"`
	Port         uint64 `toml:"port"`
	ReadTimeOut  string `toml:"read"`
	WriteTimeOut string `toml:"write"`
	IdleTimeOut  string `toml:"idle"`
}

func New() Config {
	var cfg Config

	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		loger.New().Error().Err(err)
	}

	return cfg
}
