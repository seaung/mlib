package mlib

import (
	"testing"
)

func TestMLogger(t *testing.T) {
	t.Logf("=======================\n")
	Debugf("[%s]", "debug")
	Infof("[%s]", "info")
	Warningf("[%s]", "warning")
	Errorf("[%s]", "error")
	t.Logf("=======================\n")
}
