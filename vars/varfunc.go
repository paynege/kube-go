package vars

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"paynege/kube-go/action"
)

func (hConfig *Config) ReadYaml(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, hConfig)
	return hConfig, err
}

func (app *App) CheckStatus() (bool, error) {
	status, err := action.ExecAppStatus(app.AppName)
	if err != nil {
		return false, err
	}
	switch status {
	case "active":
		return true, nil
	case "failed":
		return false, errors.New(app.AppName + "status is failed")
	case "inactive":
		return false, errors.New(app.AppName + "status is inactive")
	default:
		return false, errors.New(app.AppName + "status is unknown")
	}
}

func (app *App) SourcePathExist() (bool, error) {
	status, err := action.ExecDirectoryExist(app.AppSourcePath)
	return status, err
}

func (app *App) TargetPathExist() (bool, error) {
	status, err := action.ExecDirectoryExist(app.AppTargetPath)
	return status, err
}
