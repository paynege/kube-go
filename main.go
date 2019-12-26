package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Node struct {
	HostName string `yaml:"name"`
	HostIP   string `yaml:"IP"`
}

type Hosts struct {
	Masters []*Node `yaml:"masters"`
	Nodes   []*Node `yaml:"nodes"`
}

type App struct {
	Name       string `yaml:"name"`
	SourcePath string `yaml:"source"`
	TargetPath string `yaml:"target"`
}

type Parameter struct {
	ParaName  string `yaml:"name"`
	ParaValue string `yaml:"value"`
}

type Config struct {
	HostInfo Hosts        `yaml:"hosts"`
	AppInfo  []*App       `yaml:"apps"`
	ParaInfo []*Parameter `yaml:"parameters"`
}

var (
	setupConfig *Config
	dataConfig  []byte
)

func init() {
	setupConfig = &Config{}
	if _, err := setupConfig.readConfig(); err != nil {
		log.Fatalln(err)
	}

	dataConfig, err := json.Marshal(setupConfig)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(dataConfig))
}

func main() {
	fmt.Println(setupConfig)
}

func (config *Config) readConfig() (*Config, error) {
	file, err := ioutil.ReadFile("cfg/configure.yaml")
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(file, config); err != nil {
		return nil, err
	}
	return config, err
}
