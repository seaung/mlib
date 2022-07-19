package mlib

import (
	"os/exec"
	"strings"

	"github.com/safchain/ethtool"
	"github.com/shirou/gopsutil/v3/mem"
	ps "github.com/shirou/gopsutil/v3/net"
)

func GetMemoryTotal() uint64 {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}

	return mem.Total
}

func GetMemoryFree() uint64 {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return mem.Free
}

func GetMemoryUsed() uint64 {
	return GetMemoryTotal() - GetMemoryFree()
}

func GetMemoryUsedPercent() float64 {
	return float64(GetMemoryUsed()) / float64(GetMemoryTotal()) * 100.0
}

func GetMemoryPercent() float64 {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return 0.0
	}

	return mem.UsedPercent
}

func GetNets() (float64, float64) {
	stat, err := ps.IOCounters(true)
	if err != nil {
		return 0.0, 0.0
	}

	return float64(stat[0].BytesSent) / 1024, float64(stat[0].BytesRecv) / 1024
}

func GetInterfaces() []string {
	var inets []string
	interfaces, err := ps.Interfaces()
	if err != nil {
		return []string{"eth0", "lo"}
	}

	for _, inter := range interfaces {
		inets = append(inets, inter.Name)
	}

	return inets
}

func GetInterfacesAndState() []map[string]interface{} {
	interfaces := make([]map[string]interface{}, 0)
	cmd := exec.Command("ls", "/sys/class/net")
	buffer, err := cmd.Output()
	if err != nil {
		return []map[string]interface{}{}
	}

	for _, device := range strings.Split(string(buffer), "\n") {
		if len(device) > 1 {
			eth, err := ethtool.NewEthtool()
			if err != nil {
				continue
			}
			defer eth.Close()

			stats, err := eth.LinkState(device)
			if err != nil {
				continue
			}
			interfaces = append(interfaces, map[string]interface{}{"name": device, "status": stats})
		}
	}
	return interfaces
}
