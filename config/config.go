package config

import (
	"io/ioutil"
	"log"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
)

// Config describes the attributes that are uses to create configuration structure
type ResourceConfig struct {
	Manager *DBManagerConfig
	Consul  *consulapi.Config
}

// ControllerConfig describes the attributes for the controller configuration
type DBManagerConfig struct {
	ConsulAddress            string        `yaml:"consul_address"`
	ConsulPort               string        `yaml:"consul_port"`
	ConsulScheme             string        `yaml:"consul_scheme"`
	ConsulCAFile             string        `yaml:"consul_ca_file"`
	ConsulCertFile           string        `yaml:"consul_cert_file"`
	ConsulKeyFile            string        `yaml:"consul_key_file"`
	ConsulInsecureSkipVerify bool          `yaml:"consul_insecure_skip_verify"`
	ConsulToken              string        `yaml:"consul_token"`
	ConsulTimeout            time.Duration `yaml:"consul_timeout"`
	ConsulBasePath           string        `yaml:"consul_base_path"`
}

func LoadConfig(configLocation string) (*ResourceConfig, error) {
	newConfig := new(ResourceConfig)

	yamlFile, err := ioutil.ReadFile(configLocation)

	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &newConfig.Manager)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return newConfig, nil
}
