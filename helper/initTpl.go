package helper

import (
	"os"
	"path/filepath"
	"paynege/kube-go/oscmd"
	"text/template"
)

func GenerateTpl(value interface{}, rsc string, dist string) error {
	_, tplName := filepath.Split(rsc)
	tpl, err := template.New(tplName).ParseFiles(rsc)
	if err != nil {
		return err
	}
	result, err := oscmd.ExecDirectoryExist(dist)
	if err != nil {
		return err
	}
	if !result {
		err = oscmd.ExecMkdir(dist)
		if err != nil {
			return err
		}
	}

	var targetPath string
	if s := dist[len(dist)-1:]; s != "/" {
		targetPath = dist + "/" + tplName
	} else {
		targetPath = dist + tplName
	}
	tplFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	err = tpl.Execute(tplFile, value)
	if err != nil {
		return err
	}
	return nil
}
