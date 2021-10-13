package main

import (
	"errors"
	"fmt"
	"time"
)

type Country string

const (
	Germany      Country = "Germany"
	UnitedStates Country = "United States"
	China        Country = "China"
)

// timeZoneID 是国家=>IANA 标准时区标识符 的键值对字典
var timeZoneID = map[Country]string{
	Germany:      "Europe/Berlin",
	UnitedStates: "America/Los_Angeles",
	China:        "Asia/Shanghai",
}

//获取 IANA 时区标识符
func (c Country) TimeZoneID() (string, error) {
	if id, ok := timeZoneID[c]; ok {
		return id, nil
	}
	return "", errors.New("timezone id not found for country")
}

// 获取tz时区标识符的格式化时间字符
func TimeIn(t time.Time, tz, format string) string {

	// https:/golang.org/pkg/time/#LoadLocation loads location on
	// 加载时区
	loc, err := time.LoadLocation(tz)
	if err != nil {
		//handle error
	}
	// 获取指定时区的格式化时间字符串
	return t.In(loc).Format(format)
}

func main() {
	// 获取美国的时区结构体
	tz, err := China.TimeZoneID()
	if err != nil {
		//handle error
	}
	//格式化成美国的时区
	usTime := TimeIn(time.Now(), tz, time.RFC3339)

	fmt.Printf("Time in %s: %s",
		China,
		usTime,
	)
}
