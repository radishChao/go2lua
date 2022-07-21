// Package env 构建打包环境
package env

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func CheckEnvGo() error {
	if execute("gofmt", "-h") != nil {
		return fmt.Errorf("请安装golang")
	}
	return nil
}

func CheckEnvDocker() error {
	if execute("docker", "--version") != nil {
		return fmt.Errorf("请安装docker")
	}
	return nil
}

func execute(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &out
	c.Stderr = &stderr
	return c.Run()
}

func GetSystem() string {
	return runtime.GOOS
}
