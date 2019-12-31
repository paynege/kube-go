package main

import (
	"log"
	"paynege/kube-go/helper"
)

func main() {
	m := []map[string]interface{}{}
	m, err := helper.ParseYaml("manifest.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	for i := range m {
		helper.ParseAction(m[i])
	}
}
