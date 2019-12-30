package oscmd

import (
	"os"
	"os/exec"
	"strings"
)

func CheckAppInstalled(appName string) (bool, error) {
	cmdString := "yum info installed " + appName
	cmd := exec.Command("sh", "-c", cmdString)
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	if len(out) != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckFileExisted(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Linux systemctl is-active
func ExecAppStatus(appName string) (string, error) {
	cmdString := "systemctl status " + appName + ` | grep Active | awk -F " " '{print $2}'`
	cmd := exec.Command("sh", "-c", cmdString)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	status := strings.Replace(string(out), "\n", "", -1)
	return status, nil
}

// Linux Chmod
func ExecChmod(filename string, arg string) error {
	cmdString := "chmod " + arg + filename
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux Copy
func ExecCopy(rsc string, dist string) error {
	cmdString := "cp " + rsc + " " + dist
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl daemon-reload
func ExecDaemonReload() error {
	cmd := exec.Command("sh", "-c", "systemctl daemon-reload")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func ExecDirectoryExist(rsc string) (bool, error) {
	cmdString := "test -e " + rsc + " && echo 1 || echo 0"
	cmd := exec.Command("sh", "-c", cmdString)
	code, err := cmd.Output()
	if err != nil {
		return false, err
	}
	result := strings.Replace(string(code), "\n", "", -1)
	switch result {
	case "0":
		return false, nil
	case "1":
		return true, nil
	default:
		return false, nil
	}
}

// Linux Mkdir
func ExecMkdir(directoryPath string) error {
	cmdString := "mkdir " + directoryPath
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func ExecMove(rsc string, dist string) error {
	cmdString := "mv -r " + rsc + " " + dist
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl enable
func ExecEnableApp(appName string) error {
	cmdString := "systemctl enable " + appName
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl restart
func ExecRestartApp(appName string) error {
	cmdString := "systemctl restart " + appName
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl start
func ExecStartApp(appName string) error {
	cmdString := "systemctl start " + appName
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
