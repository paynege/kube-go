package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"paynege/kube-go/oscmd"
)

func (hConfig *Config) readYaml(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, hConfig)
	return hConfig, err
}

func (app *App) CheckStatus() (bool, error) {
	status, err := oscmd.ExecAppStatus(app.AppName)
	if err != nil {
		return false, err
	}
	switch status {
	case "active":
		return true, nil
	case "in-active":
		return false, nil
	case "load":
		return false, nil
	default:
		return false, nil
	}
}

func (app *App) SourcePathExist() (bool, error) {
	status, err := oscmd.ExecDirectoryExist(app.AppSourcePath)
	return status, err
}

func (app *App) TargetPathExist() (bool, error) {
	status, err := oscmd.ExecDirectoryExist(app.AppTargetPath)
	return status, err
}
