package main

import (
	"C"
	"log"
	"fmt"
	//	"io/ioutil"
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
func ParseConfig(configFileName string) (rc int, config_c *C.char, errstr *C.char) {
	log.Printf("configFileName: %s", configFileName)

//	fileData, err := ioutil.ReadFile(configFileName[:len(configFileName)-1])
//	log.Printf("ioutil err: %s", err)

//	log.Printf("file data: %s", fileData)

	config_go := Config{
		IOMode : "hcl:Read",
	}

	err := hclsimple.DecodeFile(configFileName[:len(configFileName)-1], nil, &config_go)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %#v", config_go)
	return 0, C.CString(fmt.Sprintf("%v", config_go)), nil
}

func main() {}

