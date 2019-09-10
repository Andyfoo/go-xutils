// Copyright 2019 Andyfoo
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package xtime

/*
时间格式参数仿php
https://www.php.net/manual/en/function.date.php

Examples:

> 当前时间:
	xTime.Now().Time

> 指定日期:
	xTime.Date(日期参数).Time

> 指定日期和时间:
	xTime.DateTime(日期和时间参数).Time


> go 默认格式:
	xTime.DateTime("20,05,2018,10,20,00").Format("2006-01-02 15:04:05") // "2018-10-11 23:22:21"

-> php 格式:
	xTime.DateTime("20,05,2018,10,20,00").PFormat("Y-m-d H:i:s") // "2018-10-11 23:22:21"
*/
import (
	"strings"
	"time"

	"github.com/Andyfoo/go-xutils/xlog"
)

var (
	months = map[int]time.Month{
		1:  time.January,
		2:  time.February,
		3:  time.March,
		4:  time.April,
		5:  time.May,
		6:  time.June,
		7:  time.July,
		8:  time.August,
		9:  time.September,
		10: time.October,
		11: time.November,
		12: time.December,
	}
	formats = map[string]string{
		"F": "January",
		"M": "Jan",
		"m": "01",
		"n": "1",
		"Y": "2006",
		"y": "06",
		"l": "Monday", //（"L"的小写字母）
		"D": "Mon",
		"d": "02",
		"j": "2",
		"H": "15",
		"G": "3",
		"i": "04",
		"s": "05",
	}

	weeks = [][]string{
		{"星期天", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		{"周日", "周一", "周二", "周三", "周四", "周五", "周六"},
		{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		{"sun", "mon", "tue", "wed", "thu", "fri", "sat"},
	}
)

//XTime 2019-08-03 08:58:48.9967745 +0800 CST m=+0.040889401
type XTime struct {
	Time time.Time
}

// Format go语言标准时间格式
func (k *XTime) Format(format string) string {
	timeLint := k.Time

	return timeLint.Format(format)
}

/*
PFormat 返回时间字符串

Custom formatters :
	"F": Long Month,
	"M": Month,
	"m": Zero Number Month,
	"n": Number Month,
	"Y": Long Year,
	"y": Year,
	"l": Long Day,
	"D": Day,
	"d": Long Number Day,
	"j": Number Day,
	"H": Long Hour,
	"G": Hour,
	"i": Long Minute,
	"s":  Long Second,


*/
func (k *XTime) PFormat(pformat string) string {
	return k.Time.Format(PFormatConv(pformat))
}

//DateStr 返回 年-月-日
func (k *XTime) DateStr() string {
	return k.PFormat("Y-m-d")
}

//TimeStr 返回 时:分:秒
func (k *XTime) TimeStr() string {
	return k.PFormat("H:i:s")
}

//DateTimeStr 返回 年-月-日 时:分:秒
func (k *XTime) DateTimeStr() string {
	return k.PFormat("Y-m-d H:i:s")
}

//Unix 时间戳（秒）
func (k *XTime) Unix() int64 {
	return k.Time.Unix()
}

//fmt.Printf("时间戳（纳秒转换为秒）：%v;\n",time.Now().UnixNano() / 1e9)
//UnixNano 时间戳（纳秒）
func (k *XTime) UnixNano() int64 {
	return k.Time.UnixNano()
}

//UnixNano 时间戳（毫秒）
func (k *XTime) UnixMilli() int64 {
	return k.Time.UnixNano() / 1e6
}

//WeekdayStr 星期字符串
// 0 = { "星期天", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六" };
// 1 = { "周日", "周一", "周二", "周三", "周四", "周五", "周六" };
// 2 = { "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" };
// 3 = { "sun", "mon", "tue", "wed", "thu", "fri", "sat" };
func (k *XTime) WeekdayStr(t ...int) string {
	if len(t) == 0 {
		return k.Time.Weekday().String()
	}
	if t[0] >= 0 && t[0] <= 3 {
		return weeks[t[0]][int(k.Time.Weekday())]
	}
	return ""
}

//Weekday 星期数字 (Sunday 星期日 = 0, ...).
func (k *XTime) Weekday() int {
	return int(k.Time.Weekday())
}

//Now 当前时间
func Now() *XTime {
	var k XTime
	k.Time = time.Now()
	return &k
}

//PFormatConv php Format 转go Format
func PFormatConv(pformat string) string {
	for key, val := range formats {
		pformat = strings.Replace(pformat, key, val, -1)
	}
	//fmt.Println(pformat)
	return pformat
}

//CustomTime 设置时间
func CustomTime(time time.Time) *XTime {
	var k XTime
	k.Time = time
	return &k
}

//Date 设置日期
func Date(day int, month int, year int) *XTime {
	var k XTime
	k.Time = time.Date(year, months[month], day, 0, 0, 0, 0, time.UTC)
	return &k
}

//DateTime 设置日期和时间
func DateTime(day int, month int, year int, hour int, min int, sec int) *XTime {
	var k XTime
	k.Time = time.Date(year, months[month], day, hour, min, sec, 0, time.UTC)
	return &k
}

//Str2Time 设置时间
func Str2Time(str string, _format ...string) *XTime {
	var format = "Y-m-d"
	if len(_format) > 0 {
		format = _format[0]
	}
	format = PFormatConv(format)
	time, err := time.Parse(format, str)
	if err != nil {
		xlog.Error("_format is ", err)
		return nil
	}
	var k XTime
	k.Time = time
	return &k
}
