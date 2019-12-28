package test

import (
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
