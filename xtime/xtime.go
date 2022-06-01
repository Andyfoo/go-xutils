// Copyright 2019 Andyfoo
// [http://andyfoo.com][http://pslib.com]
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
	"bytes"
	"time"
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
		"D": "Mon",
		"d": "02",
		"j": "2",
		"l": "Monday", //（"L"的小写字母）
		"z": "__2",

		"F": "January",
		"m": "01",
		"M": "Jan",
		"n": "1",

		"Y": "2006",
		"y": "06",

		"a": "pm",
		"A": "PM",
		"g": "3",
		"h": "03",
		"H": "15",
		"i": "04",
		"s": "05",
		"u": ".000000",

		"O": "-0700",
		"P": "-07:00",
		"T": "MST",

		"c": "2006-01-02T15:04:05Z07:00",
		"r": "Mon, 02 Jan 2006 15:04:05 -0700",
	}

	weeks = [][]string{
		{"星期天", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		{"周日", "周一", "周二", "周三", "周四", "周五", "周六"},
		{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		{"sun", "mon", "tue", "wed", "thu", "fri", "sat"},
	}

	TT  = "2006-01-02 15:04:05"
	YMD = "2006-01-02"
	HMS = "15:04:05"
)

//字段类型
type DateField uint

const (
	DateField_YEAR DateField = iota
	DateField_MONTH
	DateField_DAY
	DateField_HOUR
	DateField_MINUTE
	DateField_SECOND
)

func (p DateField) String() string {
	switch p {
	case DateField_YEAR:
		return "YEAR"
	case DateField_MONTH:
		return "MONTH"
	case DateField_DAY:
		return "DAY"
	case DateField_HOUR:
		return "HOUR"
	case DateField_MINUTE:
		return "MINUTE"
	case DateField_SECOND:
		return "SECOND"
	default:
		return "UNKNOWN"
	}
}

//XTime 2019-08-03 08:58:48.9967745 +0800 CST m=+0.040889401
type XTime struct {
	Time time.Time
}

//Now 当前时间
func Now() *XTime {
	return &XTime{
		Time: time.Now(),
	}
}

//DateStr 返回 年-月-日
func NowDateStr() string {
	return time.Now().Format(YMD)
}
func NowUtcDateStr() string {
	return time.Now().UTC().Format(YMD)
}

//TimeStr 返回 时:分:秒
func NowTimeStr() string {
	return time.Now().Format(HMS)
}
func NowUtcTimeStr() string {
	return time.Now().UTC().Format(HMS)
}

//DateTimeStr 返回 年-月-日 时:分:秒
func NowDateTimeStr() string {
	return time.Now().Format(TT)
}
func NowUtcDateTimeStr() string {
	return time.Now().UTC().Format(TT)
}

// Format go语言标准时间格式
func (k *XTime) Format(format string) string {
	return k.Time.Format(format)
}

/*
PFormat 返回时间字符串

*/
func (k *XTime) PFormat(pformat string) string {
	return k.Time.Format(PFormatConv(pformat))
}

//DateStr 返回 年-月-日
func (k *XTime) DateStr() string {
	return k.Format(YMD)
	//return k.PFormat("Y-m-d")
}

//TimeStr 返回 时:分:秒
func (k *XTime) TimeStr() string {
	return k.Format(HMS)
	//return k.PFormat("H:i:s")
}

//DateTimeStr 返回 年-月-日 时:分:秒
func (k *XTime) DateTimeStr() string {
	return k.Format(TT)
	//return k.PFormat("Y-m-d H:i:s")
}

//当天开始时间 2006-01-02 00:00:00
func (k *XTime) DayBeginDateTimeStr() string {
	return k.Format(YMD) + " 00:00:00"
}

//当天结束时间 2006-01-02 23:59:59
func (k *XTime) DayEndDateTimeStr() string {
	return k.Format(YMD) + " 23:59:59"
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

// UTC returns t with the location set to UTC.
func (t *XTime) UTC() *XTime {
	t.Time = t.Time.UTC()
	return t
}

// Local returns t with the location set to local time.
func (t *XTime) Local() *XTime {
	t.Time = t.Time.Local()
	return t
}

//日期运算
func (t *XTime) Offset(field DateField, offset int) *XTime {
	var t1 time.Time
	switch field {
	case DateField_YEAR:
		t1 = t.Time.AddDate(offset, 0, 0)
	case DateField_MONTH:
		t1 = t.Time.AddDate(0, offset, 0)
	case DateField_DAY:
		t1 = t.Time.AddDate(0, 0, offset)
	case DateField_HOUR:
		t1 = t.Time.Add(time.Hour * time.Duration(offset))
	case DateField_MINUTE:
		t1 = t.Time.Add(time.Minute * time.Duration(offset))
	case DateField_SECOND:
		t1 = t.Time.Add(time.Second * time.Duration(offset))
	default:
		t1 = t.Time
	}

	return &XTime{
		Time: t1,
	}
}

//PFormatConv php Format 转go Format
func PFormatConv(pformat string) string {
	var format bytes.Buffer
	len := len(pformat)
	for i := 0; i < len; i++ {
		val, ok := formats[pformat[i:i+1]]
		if ok {
			format.WriteString(val)
		} else {
			format.WriteString(pformat[i : i+1])
		}

	}
	//fmt.Println(pformat)
	return format.String()
}

//CustomTime 设置时间
func CustomTime(time time.Time) *XTime {
	return &XTime{
		Time: time,
	}
}

//Date 设置日期
func Date(day int, month int, year int) *XTime {
	return &XTime{
		Time: time.Date(year, months[month], day, 0, 0, 0, 0, time.UTC),
	}
}

//DateTime 设置日期和时间
func DateTime(day int, month int, year int, hour int, min int, sec int) *XTime {
	return &XTime{
		Time: time.Date(year, months[month], day, hour, min, sec, 0, time.UTC),
	}
}

//Str2Time 设置时间
func Str2Time(str string, _format ...string) *XTime {
	var format = ""
	if len(_format) > 0 {
		format = _format[0]
	} else if len(_format) == 0 && len(str) == 19 {
		format = "Y-m-d H:i:s"
	} else if len(_format) == 0 && len(str) == 10 {
		format = "Y-m-d"
	} else {
		format = "Y-m-d"
	}
	format = PFormatConv(format)
	time, err := time.Parse(format, str)
	if err != nil {
		//xlog.Error("_format is ", err)
		return nil
	}
	return &XTime{
		Time: time,
	}
}
