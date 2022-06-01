package xtime

import (
	"fmt"
	"testing"
	"time"
)

// func TestSpeed1(t *testing.T) {
// 	for i := 0; i < 10000000; i++ {
// 		time.Now().Format("2006-01-02 15:04:05")
// 	}
// }
// func TestSpeed2(t *testing.T) {
// 	for i := 0; i < 10000000; i++ {
// 		Now().PFormat("Y-m-d H:i:s")
// 		//NowDateTimeStr()
// 	}
// }

func Test1(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05.999999999Z07:00 2 __2 002"))
	fmt.Println(time.Now().Format("Z07:00:00"))
	fmt.Println(time.Unix(1588009973, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(">>>", PFormatConv("Y-m-d H:i:su"))
	fmt.Println(Now().PFormat("Y-m-d H:i:su"))
	fmt.Println(Now().Unix())
	fmt.Println(Now().UnixMilli())
	fmt.Println(Now().UnixNano())
	fmt.Println(Now().Weekday())
	fmt.Println(Now().WeekdayStr(1))
	fmt.Println(Str2Time("2018-04-23 23:11:23", "Y-m-d H:i:s").PFormat("Y-m-d H:i:s"))

}

func Test2(t *testing.T) {
	fmt.Println(DateField_DAY.String())

	fmt.Println(Now().Offset(DateField_YEAR, 3).UTC().DateTimeStr())
	fmt.Println(Now().Offset(DateField_MONTH, 3).PFormat("Y-m-d H:i:s"))
	fmt.Println(Now().Offset(DateField_DAY, 3).PFormat("Y-m-d H:i:s"))

	fmt.Println(Now().Offset(DateField_DAY, -113).PFormat("Y-m-d H:i:s"))

	fmt.Println(Now().Offset(DateField_HOUR, 3).PFormat("Y-m-d H:i:s"))
	fmt.Println(Now().Offset(DateField_MINUTE, 3).PFormat("Y-m-d H:i:s"))
	fmt.Println(Now().Offset(DateField_SECOND, 3).PFormat("Y-m-d H:i:s"))

	fmt.Println(Now().DayBeginDateTimeStr(), Now().DayEndDateTimeStr())
}

func TestDay(t *testing.T) {
	/*
	   d  月份中的第几天，有前导零的 2 位数字 01 到 31
	   D  星期中的第几天，文本表示，3 个字母 Mon 到 Sun
	   j  月份中的第几天，没有前导零 1 到 31
	   l（“L”的小写字母） 星期几，完整的文本格式 Sunday 到 Saturday
	   ---不支持---N  ISO-8601 格式数字表示的星期中的第几天 1（表示星期一）到 7（表示星期天）
	   ---不支持---S  每月天数后面的英文后缀，2 个字符 st，nd，rd 或者 th。可以和 j 一起用
	   ---不支持---w  星期中的第几天，数字表示 0（表示星期天）到 6（表示星期六）
	   z  年份中的第几天 0 到 365

	   星期  --- ---
	   ---不支持---W  ISO-8601 格式年份中的第几周，每周从星期一开始 例如：42（当年的第 42 周）
	*/
	out("d")
	out("D")
	out("j")
	out("l")
	out("z Y-m-d H:i:s P")
	outUTC("z Y-m-d H:i:s P")
	out("Y-m-d H:i:s P")
}
func TestMonth(t *testing.T) {
	/*
		月  --- ---
		F  月份，完整的文本格式，例如 January 或者 March January 到 December
		m  数字表示的月份，有前导零 01 到 12
		M  三个字母缩写表示的月份 Jan 到 Dec
		n  数字表示的月份，没有前导零 1 到 12
		---不支持---t  给定月份所应有的天数 28 到 31
	*/
	out("F")
	out("m")
	out("M")
	out("n")
}
func TestYear(t *testing.T) {
	/*
		年  --- ---
		---不支持---L  是否为闰年 如果是闰年为 1，否则为 0
		---不支持---o  ISO-8601 格式年份数字。这和 Y 的值相同，只除了如果 ISO 的星期数（W）属于前一年或下一年，则用那一年。 Examples: 1999 or 2003
		Y  4 位数字完整表示的年份 例如：1999 或 2003
		y  2 位数字表示的年份 例如：99 或 03
	*/
	out("Y")
	out("y")
}
func TestTime(t *testing.T) {
	/*
		时间  --- ---
		a  小写的上午和下午值 am 或 pm
		A  大写的上午和下午值 AM 或 PM
		---不支持---B  Swatch Internet 标准时 000 到 999
		g  小时，12 小时格式，没有前导零 1 到 12
		---不支持---G  小时，24 小时格式，没有前导零 0 到 23
		h  小时，12 小时格式，有前导零 01 到 12
		H  小时，24 小时格式，有前导零 00 到 23
		i  有前导零的分钟数 00 到 59>
		s  秒数，有前导零 00 到 59>
		u  毫秒 .000000
	*/
	out("a")
	out("A")
	out("g")
	out("h")
	out("H")
	out("i")
	out("s")
	out("u")
}
func TestTimeZone(t *testing.T) {
	/*
		时区  --- ---
		---不支持---e  时区标识 例如：UTC，GMT，Atlantic/Azores
		---不支持---I  是否为夏令时 如果是夏令时为 1，否则为 0
		O  与格林威治时间相差的小时数 例如：+0200
		P  与格林威治时间（GMT）的差别，小时和分钟之间有冒号分隔 例如：+02:00
		T  本机所在的时区 例如：EST，MDT
		---不支持---Z  时差偏移量的秒数。UTC 西边的时区偏移量总是负的，UTC 东边的时区偏移量总是正的。 -43200 到 43200
	*/
	out("O")
	out("P")
	out("T")
}
func TestFullTime(t *testing.T) {
	/*
		完整的日期／时间  --- ---
		c  ISO 8601 格式的日期 2004-02-12T15:19:21+00:00
		r  RFC 822 格式的日期 例如：Thu, 21 Dec 2000 16:01:07 +0200
		---不支持---U  从 Unix 纪元（January 1 1970 00:00:00 GMT）开始至今的秒数
	*/
	out("c")
	out("r")
}
func out(dateFormat string) {
	fmt.Printf("%s = (format:%s, value:%s)\n", dateFormat, PFormatConv(dateFormat), Now().PFormat(dateFormat))
}
func outUTC(dateFormat string) {
	fmt.Printf("%s = (format:%s, value:%s)\n", dateFormat, PFormatConv(dateFormat), Now().UTC().PFormat(dateFormat))
}
