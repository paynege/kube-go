package test

import (
	"paynege/kube-go/oscmd"
	"testing"
)

func TestExecAppStatus(t *testing.T) {
	status, _ := oscmd.ExecAppStatus("sshd")
	if status != "active" {
		t.Error(`ExecAppStatus("sshd") is ` + status + `, expect "active"`)
	}
	status, _ = oscmd.ExecAppStatus("docker")
	if status != "loaded" {
		t.Error(`ExecAppStatus("docker") is ` + status + `, expect "loaded"`)
	}
}
