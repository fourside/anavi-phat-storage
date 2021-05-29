package main

import (
	"time"
)

func parseTime(utcDateString string) (string, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	utcTime, err := time.Parse("20060102-150405", utcDateString)
	jstTime := utcTime.In(jst)
	if err != nil {
		return "", err
	}
	return jstTime.Format(time.RFC3339), nil
}
