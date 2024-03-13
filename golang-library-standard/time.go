package main

import (
	"fmt"
	"time"
)

func main() {
	// untuk check waktu sekarang
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	// membuat waktu
	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// parsing waktu dengan RFC3339
	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"
	// value := "ASAL"

	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println(valueTime.Minute())
}
