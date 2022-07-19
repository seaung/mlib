package mlib

import (
	"testing"
)

func TestMTime(t *testing.T) {
	t.Log("=========================")
	t.Logf("get current timestamp : %d\n", GetCurrentTimestampMis())
	t.Logf("get current timestamps : %d\n", GetCurrentTimestampNs())
	t.Logf("get time to string : %s\n", TimetoString(182239966986))
	t.Logf("get timestapm to time : %s\n", Timestamp2Time(165593434322))
	t.Logf("get string to time : %v\n", TimeStringToTime("2022-08-06 12:33:08"))
	t.Log("=========================")
}
