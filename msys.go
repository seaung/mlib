package mlib

import (
	"bytes"
	"net"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func ExecuteCmd(cmd string, options ...string) string {
	c := exec.Command(cmd, options...)

	output, err := c.CombinedOutput()
	if err != nil {
		return ""
	}

	return string(output)
}

func ExecuteCmdWithRun(cmd string, options ...string) string {
	c := exec.Command(cmd, options...)

	var stdout bytes.Buffer
	c.Stdout = &stdout

	err := c.Run()
	if err != nil {
		return ""
	}

	return stdout.String()
}

func GetMacAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "Unknow Mac"
	}
	return interfaces[0].HardwareAddr.String()
}

func GetMacAddresses() []string {
	var macs []string
	interfaces, err := net.Interfaces()
	if err != nil {
		return macs
	}

	for _, v := range interfaces {
		macs = append(macs, v.HardwareAddr.String())
	}
	return macs
}

func GetLocalIP() string {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, addr := range address {
		if ipNet, isIP := addr.(*net.IPNet); isIP && !ipNet.IP.IsLoopback() {
			return ipNet.String()
		}
	}
	return "localhost"
}

func GetInternalIP() string {
	internals, err := net.Interfaces()
	if err != nil {
		return "localhost"
	}

	for _, inet := range internals {
		if inet.Flags&net.FlagUp != 0 && !strings.HasPrefix(inet.Name, "lo") {
			addrs, err := inet.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return "localhost"
}

func GetCurrentUser() string {
	u, err := user.Current()
	if err != nil {
		return "Unknow"
	}
	return u.Username
}

func GetSystemArch() string {
	return runtime.GOARCH
}

func GetSystemType() string {
	return runtime.GOOS
}

func GetComputerName() string {
	name, err := os.Hostname()
	if err != nil {
		return "Unknow"
	}
	return name
}
