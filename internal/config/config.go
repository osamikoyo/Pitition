package config

import (
	"github.com/BurntSushi/toml"
	"github.com/osamikoyo/pitition/pkg/loger"
)

type Config struct {
	Hostname     string `toml:"hostname"`
	Port         uint64 `toml:"port"`
	ReadTimeOut  uint16 `toml:"read"`
	WriteTimeOut uint16 `toml:"write"`
	IdleTimeOut  uint16 `toml:"idle"`
}

func New() Config {
	var cfg Config

	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		loger.New().Error().Err(err)
	}

	return cfg
}
