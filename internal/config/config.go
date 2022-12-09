package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		Addr    string
		Bridges []Bridge
	}
	Bridge struct {
		Enable   *bool
		Name     string
		Addr     string
		Patterns []string
	}
)

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = toml.Unmarshal(data, &conf)
	return &conf, err
}

func (b Bridge) Key() string {
	return b.Addr
}
