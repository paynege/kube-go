package test

import (
	"paynege/kube-go/oscmd"
	"strings"
	"testing"
)

func TestExecAppStatus(t *testing.T) {
	status, _ := oscmd.ExecAppStatus("sshd")
	if !strings.EqualFold(status, "active") {
		t.Error(`ExecAppStatus("sshd") is ` + status + `, expect "active"`)
	}
	status, _ = oscmd.ExecAppStatus("docker")
	if !strings.EqualFold(status, "inactive") {
		t.Error(`ExecAppStatus("docker") is ` + status + `, expect "inactive"`)
	}
}
