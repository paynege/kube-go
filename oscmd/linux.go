package oscmd

import (
	"bufio"
	"io"
	"os/exec"
)

// Linux systemctl is-active
func ExecAppStatus(appName string) (string, error) {
	cmdstring := "systemctl is-active " + appName
	cmd := exec.Command("bash", "-c", cmdstring)
	code, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(code)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			return line, err
		}
	}
}

// Linux Chmod
func ExecChmod(filename string, arg string) error {
	cmd := exec.Command("chmod", arg, filename)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux Copy
func ExecCopy(rsc string, dist string) error {
	cmd := exec.Command("cp", "-r", rsc, dist)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl daemon-reload
func ExecDaemonReload() error {
	cmd := exec.Command("systemctl", "daemon-reload")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func ExecDirectoryExist(rsc string) (bool, error) {
	cmd := exec.Command("test", "-e", "rsc", "&& echo 1 || echo 0")
	code, err := cmd.StdoutPipe()
	if err != nil {
		return false, err
	}
	reader := bufio.NewReader(code)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			switch line {
			case "0":
				return false, nil
			case "1":
				return true, nil
			default:
				return false, nil
			}
		}
	}
}

// Linux Mkdir
func ExecMkdir(directoryPath string) error {
	cmd := exec.Command("mkdir", "-r", directoryPath)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func ExecMove(rsc string, dist string) error {
	cmd := exec.Command("mv", "-r", rsc, dist)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl enable
func ExecEnableApp(appName string) error {
	cmd := exec.Command("systemctl", "enable", appName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl restart
func ExecRestartApp(appName string) error {
	cmd := exec.Command("systemctl", "restart", appName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Linux systemctl start
func ExecStartApp(appName string) error {
	cmd := exec.Command("systemctl", "start", appName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
