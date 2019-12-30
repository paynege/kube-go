package test

import (
	"paynege/kube-go/action"
	"testing"
)

func TestExecAppStatus(t *testing.T) {
	if status, err := action.ServiceStatus("firewalld"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if status != "active" {
		t.Error(`ExecAppStatus("firewalld") is ` + status + `, expect "active"`)
	}

	if status, err := action.ServiceStatus("docker"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if status != "active" {
		t.Error(`ExecAppStatus("docker") is ` + status + `, expect "active"`)
	}

	if status, err := action.ServiceStatus("sshd"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if status != "active" {
		t.Error(`ExecAppStatus("sshd") is ` + status + `, expect "active"`)
	}
}

func TestExecDirectoryExist(t *testing.T) {
	if bool, err := action.IfDirectoryExisted("../template"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if bool != true {
		t.Error("template does not exist, expect existed")
	}
	if bool, err := action.IfDirectoryExisted("../root"); err != nil {
		t.Error(`Error occur: ` + err.Error())
	} else if bool != false {
		t.Error("root existed, expect not existed")
	}
}

func TestGenerateTpl(t *testing.T) {
	etcd_dist := "etcd_ca"
	k8s_dist := "k8s_ca"
	err := action.GenerateTpl("87600h", "../template/etcd/ca-config.json", etcd_dist)
	if err != nil {
		t.Error(`Error occur: ` + err.Error())
	}
	err = action.GenerateTpl("192.168.3.100", "../template/kubernetes/admin-csr.json", k8s_dist)
	if err != nil {
		t.Error(`Error occur: ` + err.Error())
	}
}

/*

func TestGenerateCert(t *testing.T) {
	etcdDist := "etcd_ca"
	//k8sDist := "k8s_ca"
	err := helper.GenerateCertFile("etcd", etcdDist)
	if err != nil {
		t.Error(err.Error())
	} else {
		if status, err := action.CheckFileExisted("etcd_ca/ca.pem"); status == false && err == nil {
			t.Error("ca.pem not generated")
		}
		if status, err := action.CheckFileExisted("etcd_ca/ca.csr"); status != true && err == nil {
			t.Error("ca.csr not generated")
		}
		if status, err := action.CheckFileExisted("etcd_ca/ca-key.pem"); status != true && err == nil {
			t.Error("etcd_ca/ca-key.pem not generated")
		}
		if status, err := action.CheckFileExisted("etcd_ca/server.pem"); status != true && err == nil {
			t.Error("etcd_ca/server.pem not generated")
		}
		if status, err := action.CheckFileExisted("etcd_ca/server-key.pem"); status != true && err == nil {
			t.Error("etcd_ca/server-key.pem not generated")
		}
	}
}

*/
