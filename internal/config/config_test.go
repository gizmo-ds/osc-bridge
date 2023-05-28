package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromFile(t *testing.T) {
	conf, err := LoadFromFile("../../config.toml")
	assert.NoError(t, err)
	assert.Equal(t, conf.Addr, "127.0.0.1:9001")
	assert.Equal(t, len(conf.Bridges), 2)
	assert.NotNil(t, conf.Bridges[0])
	assert.Equal(t, *conf.Bridges[0].Enable, false)
	assert.Nil(t, conf.Bridges[1].Enable)
}

func TestLoad(t *testing.T) {
	testConfig := `addr = "127.0.0.1:9001"

	[[bridges]]
	enable = false
	name = "B1"
	addr = "127.0.0.1:10001"
	patterns = [
			"/avatar/change",
	]
	
	[[bridges]]
	name = "B2"
	addr = "127.0.0.1:10002"
	patterns = [
			"/avatar/change",
	]
	`
	conf, err := Load([]byte(testConfig))
	assert.NoError(t, err)
	assert.Equal(t, conf.Addr, "127.0.0.1:9001")
	assert.Equal(t, len(conf.Bridges), 2)
	assert.NotNil(t, conf.Bridges[0])
	assert.Equal(t, *conf.Bridges[0].Enable, false)
	assert.Nil(t, conf.Bridges[1].Enable)
}
