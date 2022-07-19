package mlib

import (
	"testing"
)

func TestSys(t *testing.T) {
	sout := ExecuteCmd("ls", "-al")
	t.Log(sout)

	t.Log("============================")

	sout = ExecuteCmdWithRun("ls", "-al")
	t.Log(sout)

	t.Log("============================")

	t.Logf("get current user : %s", GetCurrentUser())
	t.Logf("get system arch : %s", GetSystemArch())
	t.Logf("get system type : %s", GetSystemType())
	t.Logf("get computer name : %s", GetComputerName())
	t.Logf("get local ip : %s", GetLocalIP())
	t.Logf("get internal ip : %s", GetInternalIP())
	t.Logf("get mac : %s", GetMacAddress())
	t.Logf("get mac list : %s", GetMacAddresses())
}
