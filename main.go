package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Node struct {
	NodeName string `yaml:"name"`
	NodeIP   string `yaml:"IP"`
}

type Host struct {
	Master []*Node `yaml:"masters"`
	Nodes  []*Node `yaml:"nodes"`
}

type Config struct {
	HostInfo Host         `yaml:"hosts"`
	AppInfo  []*App       `yaml:"apps"`
	ParaInfo []*Parameter `yaml:"parameters"`
}

type App struct {
	AppName       string `yaml:"name"`
	AppSourcePath string `yaml:"source"`
	AppTargetPath string `yaml:"target"`
}

type Parameter struct {
	ParaName  string `yaml:"name"`
	ParaValue string `yaml:"value"`
}

var (
	hostConfig Config
)

func (hConfig *Config) readYaml(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, hConfig)
	return hConfig, err
}

func main() {
	hostConfig = Config{}

	if _, err := hostConfig.readYaml("cfg/configure.yaml"); err != nil {
		log.Fatalln(err)
	}

}
