package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	var cfg *ResourceConfig
	var err error
	t.Parallel()

	cfg, err = LoadConfig("../examples/config.yaml")

	assert.Nil(t, err, "There was not issues in the call")

	assert.Equal(t, cfg.Manager.ConsulScheme, "http", "wrong Schema")
	assert.Equal(t, cfg.Manager.ConsulAddress, "localhost", "wrong Address")
	assert.Equal(t, cfg.Manager.ConsulPort, "8500", "wrong Port")
	assert.Equal(t, cfg.Manager.ConsulToken, "token", "wrong Token")

}
