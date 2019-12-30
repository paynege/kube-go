package helper

import (
	"github.com/pkg/errors"
	"os/exec"
	"path/filepath"
)

func GenerateCertFile(appName string, jsonPath string) error {
	switch appName {
	case "etcd":
		err := generateETCDKey(jsonPath)
		return err
	case "kubernetes":
		err := generateK8SKeys(jsonPath)
		return err
	default:
		return errors.New("Not supported App")
	}
}

func initializeCA(jsonPath string) error {
	filePath, fileName := filepath.Split(jsonPath)
	cmdString := "cfssl gencert -initca " + fileName + " | cfssljson -bare ca -"
	cmd := exec.Command("sh", "-c", cmdString)
	cmd.Dir = filePath
	err := cmd.Run()
	return err
}

func generateKey(jsonPath string, appName string, outer string) error {
	cmdString := "cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=" + appName + " " + outer + "-csr.json | cfssljson -bare " + outer
	cmd := exec.Command("sh", "-c", cmdString)
	cmd.Dir = jsonPath
	err := cmd.Run()
	return err
}

func generateETCDKey(jsonPath string) error {
	err := initializeCA(jsonPath)
	if err != nil {
		return err
	}
	err = generateKey(jsonPath, "etcd", "server")
	return err
}

func generateK8SKeys(jsonPath string) error {
	err := initializeCA(jsonPath)
	if err != nil {
		return err
	}
	err = generateKey(jsonPath, "kubernetes", "apiserver")
	if err != nil {
		return err
	}
	err = generateKey(jsonPath, "kubernetes", "kube-proxy")
	if err != nil {
		return err
	}
	err = generateKey(jsonPath, "kubernetes", "admin")
	if err != nil {
		return err
	}
	return nil
}
