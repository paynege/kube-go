package action

import (
	"os"
	"os/exec"
	"strings"
)

func IfAppInstalled(appName string) (bool, error) {
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

func IfFileExisted(file string) (bool, error) {
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
func ServiceStatus(serviceName string) (string, error) {
	cmdString := "systemctl status " + serviceName + ` | grep Active | awk -F " " '{print $2}'`
	cmd := exec.Command("sh", "-c", cmdString)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	status := strings.Replace(string(out), "\n", "", -1)
	return status, nil
}

// Linux Chmod
func Chmod(file string, arg string) error {
	cmdString := "chmod " + arg + file
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux Copy
func Copy(rsc string, dist string) error {
	cmdString := "cp " + rsc + " " + dist
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl daemon-reload
func DaemonReload() error {
	cmd := exec.Command("sh", "-c", "systemctl daemon-reload")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func IfDirectoryExisted(directory string) (bool, error) {
	cmdString := "test -e " + directory + " && echo 1 || echo 0"
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
func Mkdir(directory string) error {
	cmdString := "mkdir " + directory
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func Move(rsc string, dist string) error {
	cmdString := "mv -r " + rsc + " " + dist
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl enable
func ServiceEnable(serviceName string) error {
	cmdString := "systemctl enable " + serviceName
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl restart
func ServiceRestart(serviceName string) error {
	cmdString := "systemctl restart " + serviceName
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl start
func ServiceStart(serviceName string) error {
	cmdString := "systemctl start " + serviceName
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl stop
func ServiceStop(serviceName string) error {
	cmdString := "systemctl stop " + serviceName
	cmd := exec.Command("sh", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func Reboot() error {
	cmd := exec.Command("sh", "-c", "reboot")
	err := cmd.Run()
	return err
}

func PowerOff() error {
	cmd := exec.Command("sh", "-c", "poweroff")
	err := cmd.Run()
	return err
}
