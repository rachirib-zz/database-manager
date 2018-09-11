package consul

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
	"strings"

	"ac2.io/hc/database-manager/config"
	consulapi "github.com/hashicorp/consul/api"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// Adapter builds configuration and returns Consul Client
type Adapter struct {
	client *consulapi.Client
	Config *consulapi.Config
}

// New returns the ConsulAdapter.
func (c *Adapter) New(cfg *config.ResourceConfig) *Adapter {
	var address string
	var err error
	var uri *url.URL
	var tlsConfig *tls.Config

	//Build URI
	address = fmt.Sprintf("%s://%s:%s",
		cfg.Manager.ConsulScheme, cfg.Manager.ConsulAddress, cfg.Manager.ConsulPort)

	uri, err = url.Parse(address)
	if err != nil {
		log.Fatalf("bad adapter uri: ")
	}

	switch uri.Scheme {
	case "consul-unix":
		cfg.Consul.Address = strings.TrimPrefix(uri.String(), "consul-")

	case "https":
		tlsConfigDesc := &consulapi.TLSConfig{
			Address:            uri.Host,
			CAFile:             cfg.Manager.ConsulCAFile,
			CertFile:           cfg.Manager.ConsulCertFile,
			KeyFile:            cfg.Manager.ConsulKeyFile,
			InsecureSkipVerify: cfg.Manager.ConsulInsecureSkipVerify,
		}
		tlsConfig, err = consulapi.SetupTLSConfig(tlsConfigDesc)
		if err != nil {
			log.Fatalf("Cannot set up Consul TLSConfig: %s", err)
		}
		cfg.Consul.Scheme = uri.Scheme
		transport := cleanhttp.DefaultPooledTransport()
		transport.TLSClientConfig = tlsConfig
		cfg.Consul.Transport = transport
		cfg.Consul.Address = uri.Host

	default:
		cfg.Consul.Address = uri.Host
	}

	// Add Token
	if cfg.Manager.ConsulToken != "" {
		cfg.Consul.Token = cfg.Manager.ConsulToken
	}

	//Timeout
	cfg.Consul.Transport.IdleConnTimeout = cfg.Manager.ConsulTimeout

	client, err := consulapi.NewClient(cfg.Consul)
	if err != nil {
		log.Fatalf("consul: %s", uri.Scheme)
	}

	//Store config
	c.Config = cfg.Consul

	c.client = client
	return c
}

// Retrieve all the consul keys
func (c *Adapter) KVs(prefix string) []string {
	keys, _, err := c.client.KV().Keys(prefix, "", nil)
	if err != nil {
		log.Fatal("Error retrieving consul keys list")
	}

	return keys
}
