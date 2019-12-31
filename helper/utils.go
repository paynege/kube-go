package helper

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"paynege/kube-go/action"
)

func ParseYaml(file string) ([]map[string]interface{}, error) {
	var m []map[string]interface{}
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &m)
	return m, nil
}

func ParseAction(a map[string]interface{}) error {
	switch a["name"] {
	case "copy":
		err := action.Copy(a["src"].(string), a["dest"].(string))
		return err

	case "chmod":
		var p string
		if v, ok := a["exec"]; ok {
			if v.(bool) {
				p += " +x"
			} else {
				p += " -x"
			}
		}
		if v, ok := a["read"]; ok {
			if v.(bool) {
				p += " +r"
			} else {
				p += " -r"
			}
		}
		if v, ok := a["write"]; ok {
			if v.(bool) {
				p += " +w"
			} else {
				p += " -w"
			}
		}
		err := action.Chmod(a["file"].(string), p)
		return err
	default:
		return errors.New("Unsupported action")
	}
}
