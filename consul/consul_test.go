package consul

import (
	"testing"
	"time"

	"ac2.io/hc/database-manager/config"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	cfg := &config.ResourceConfig{
		Manager: &config.DBManagerConfig{
			ConsulAddress: "localhost",
			ConsulPort:    "8500",
			ConsulScheme:  "http",
			ConsulToken:   "token",
			ConsulTimeout: time.Duration(0),
		},
		Consul: consulapi.DefaultConfig(),
	}

	consulInstance := Adapter{}
	consulInstance.New(cfg)

	assert.Equal(t, cfg.Consul.Address, "localhost:8500", "wrong URI")
	assert.Equal(t, cfg.Consul.Scheme, "http", "wrong scheme")
	assert.Equal(t, cfg.Consul.Token, "token", "wrong token")
	assert.Equal(t, cfg.Consul.HttpClient.Timeout, time.Duration(0), "wrong timeout")

	assert.Equal(t, "localhost:8500", cfg.Consul.Address, "wrong URI")

}

func TestConsulAdapterMethods(t *testing.T) {
	var search []string
	t.Parallel()

	cfg := &config.ResourceConfig{
		Manager: &config.DBManagerConfig{
			ConsulAddress: "localhost",
			ConsulPort:    "8500",
			ConsulScheme:  "http",
			ConsulToken:   "token",
			ConsulTimeout: time.Duration(0),
		},
		Consul: consulapi.DefaultConfig(),
	}

	consulInstance := Adapter{}
	consulAgent := consulInstance.New(cfg)

	search = consulAgent.KVs("ac2/")
	assert.NotNil(t, search, "searching values are nil")

}
