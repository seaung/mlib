package mlib

import (
	"time"
)

func GetCurrentTimestampMis() int64 {
	return time.Now().Unix()
}

func GetCurrentTimestampNs() int64 {
	return time.Now().UnixNano()
}

func TimetoString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func Timestamp2Time(tm int64) time.Time {
	tunix := TimetoString(tm)
	return TimeStringToTime(tunix)
}

func TimeStringToTime(timeString string) time.Time {
	fmtime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return time.Now()
	}
	return fmtime
}
