package main

import (
	"log"
	"paynege/kube-go/vars"
)

func main() {
	hostConfig := vars.Config{}

	if _, err := hostConfig.ReadYaml("cfg/configure.yaml"); err != nil {
		log.Fatalln(err)
	}

	//hostConfig.AppInfo[0].CheckStatus()
}
