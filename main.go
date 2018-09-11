package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"ac2.io/hc/database-manager/config"
	"ac2.io/hc/database-manager/serve"
)

var (
	// VERSION is filled out during the build process (using git describe output)
	VERSION = "0.0.1"

	cfg         *config.ResourceConfig
	configMap   = flag.String("config", "/tmp/config.yaml", "location of the Config that containes the custom configuration to use")
	versionFlag = flag.Bool("version", false, "print version end exit")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	cfg, err := config.LoadConfig(*configMap)
	if err != nil {
		log.Printf("Not able to load the configuration   #%v ", err)
	}

	log.Print("Configuration Loaded")
	log.Printf("Consul Address: %q", cfg.Manager.ConsulAddress)
	log.Printf("Consul Port: %q", cfg.Manager.ConsulPort)
	log.Printf("Consul Scheme: %q", cfg.Manager.ConsulScheme)
	log.Printf("Consul Base Path: %q", cfg.Manager.ConsulBasePath)

	router := serve.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
