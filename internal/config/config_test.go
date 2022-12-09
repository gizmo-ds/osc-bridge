package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	conf, err := Load("config.toml")
	assert.NoError(t, err)
	assert.Equal(t, conf.Addr, "127.0.0.1:9001")
	assert.Equal(t, len(conf.Bridges), 2)
	assert.NotNil(t, conf.Bridges[0])
	assert.Equal(t, *conf.Bridges[0].Enable, false)
	assert.Nil(t, conf.Bridges[1].Enable)
}
