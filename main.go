package main

import (
	"log"
)

func main() {
	hostConfig = Config{}

	if _, err := hostConfig.readYaml("cfg/configure.yaml"); err != nil {
		log.Fatalln(err)
	}

	/*
		for _, v := range hostConfig.AppInfo {
			if err := oscmd.ExecMkdir(v.AppSourcePath); err != nil {
				log.Println(err)
			}
			if err := oscmd.ExecMkdir(v.AppTargetPath); err != nil {
				log.Println(err)
			}
		}

	*/
	//hostConfig.AppInfo[0].CheckStatus()
}
