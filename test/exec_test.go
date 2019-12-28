package test

import (
	"paynege/kube-go/helper"
	//"paynege/kube-go/oscmd"
	"testing"
)

/*
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
*/

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
