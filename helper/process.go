package helper

import (
	"log"
	"paynege/kube-go/oscmd"
	"paynege/kube-go/vars"
)

func PrepareDirectory(appConfig *vars.Config) {
	for _, v := range appConfig.AppInfo {
		if err := oscmd.ExecMkdir(v.AppSourcePath); err != nil {
			log.Println(err)
		}
		if err := oscmd.ExecMkdir(v.AppTargetPath); err != nil {
			log.Println(err)
		}
	}
	//if err := oscmd.ExecMkdir("")
}
