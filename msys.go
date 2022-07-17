package mlib

import (
	"net"
	"os"
	"os/user"
	"runtime"
)

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
