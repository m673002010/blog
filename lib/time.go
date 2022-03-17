package lib

import (
	"time"
)

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetDay() string {
	tmpl := "20060102"
	return time.Now().Format(tmpl)
}

func GetDate() string {
	tmpl := "2021-01-01 15:05:05"
	return time.Now().Format(tmpl)
}
