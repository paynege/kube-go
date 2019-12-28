package test

import (
	"paynege/kube-go/helper"
	"paynege/kube-go/oscmd"
	"testing"
)

func TestExecAppStatus(t *testing.T) {
	if status, err := oscmd.ExecAppStatus("firewalld"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if status != "active" {
		t.Error(`ExecAppStatus("firewalld") is ` + status + `, expect "active"`)
	}

	if status, err := oscmd.ExecAppStatus("docker"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if status != "active" {
		t.Error(`ExecAppStatus("docker") is ` + status + `, expect "active"`)
	}

	if status, err := oscmd.ExecAppStatus("sshd"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if status != "active" {
		t.Error(`ExecAppStatus("sshd") is ` + status + `, expect "active"`)
	}
}

func TestExecDirectoryExist(t *testing.T) {
	if bool, err := oscmd.ExecDirectoryExist("../template"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if bool != true {
		t.Error("template does not exist, expect existed")
	}
	if bool, err := oscmd.ExecDirectoryExist("../root"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if bool != false {
		t.Error("root existed, expect not existed")
	}
}

func TestGenerateTpl(t *testing.T) {
	etcd_dist := "etcd_ca"
	k8s_dist := "k8s_ca"
	err := helper.GenerateTpl("87600h", "../template/etcd/ca-config.json", etcd_dist)
	if err != nil {
		t.Error(`Error occur: ` + err.Error())
	}
	err = helper.GenerateTpl("192.168.3.100", "../template/kubernetes/admin-csr.json", k8s_dist)
	if err != nil {
		t.Error(`Error occur: ` + err.Error())
	}
}
