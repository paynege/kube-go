package test

import (
	"paynege/kube-go/oscmd"
	"testing"
)

func TestExecAppStatus(t *testing.T) {
	if status, _ := oscmd.ExecAppStatus("sshd"); status != "active" {
		t.Error(`ExecAppStatus("sshd") is ` + status + `, expect "active"`)
	}
}

func TestExecAppStatus1(t *testing.T) {
	if status, _ := oscmd.ExecAppStatus("firewalld"); status != "inactive" {
		t.Error(`ExecAppStatus("firewalld") is ` + status + `, expect "inactive"`)
	}
}

func TestExecAppStatus2(t *testing.T) {
	if status, _ := oscmd.ExecAppStatus("docker"); status != "inactive" {
		t.Error(`ExecAppStatus("docker") is ` + status + `, expect "inactive"`)
	}
}
