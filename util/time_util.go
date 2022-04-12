package util

import (
	"fmt"
	"time"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

type Time time.Time

func GetDateNow() string {
	return time.Now().Format(DateFormat)
}

func GetTimeNow() string {
	return time.Now().Format(TimeFormat)
}

func GetLocationTime(timeStr string) time.Time {
	location, err := time.LoadLocation("Asia/Taipei")
	locTime, err := time.ParseInLocation(TimeFormat, timeStr, location)

	if err != nil {
		locTime = time.Now()
		fmt.Println("error:", err)
	}
	return locTime
}
