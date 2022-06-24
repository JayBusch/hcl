package main

import (
	"C"
	"log"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

	type Config struct {
	IOMode  string        `hcl:"io_mode"`
	Service ServiceConfig `hcl:"service,block"`
}

type ServiceConfig struct {
	Protocol   string          `hcl:"protocol,label"`
	Type       string          `hcl:"type,label"`
	ListenAddr string          `hcl:"listen_addr"`
	Processes  []ProcessConfig `hcl:"process,block"`
}

type ProcessConfig struct {
	Type    string   `hcl:"type,label"`
	Command []string `hcl:"command"`
}

//export ParseConfig
func ParseConfig() (rc int, return_config *C.char, errstr *C.char) {
	var config Config
	err := hclsimple.DecodeFile("config.hcl", nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %#v", config)
	return 0, C.CString(fmt.Sprintf("%v", config)), nil
}

func main() {}

