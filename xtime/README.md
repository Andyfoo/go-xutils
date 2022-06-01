```go
func Test1(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(1588009973, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(">>>", PFormatConv("Y-m-d H:i:s.SSS"))
	fmt.Println(Now().PFormat("Y-m-d H:i:s.SSS"))
	fmt.Println(Now().PFormat("Y-m-d H:i:s.SSSSSSSSS"))
	fmt.Println(Now().PFormat("Y-m-d H:i:s.SSSSSS"))
	fmt.Println(Now().Unix())
	fmt.Println(Now().UnixMilli())
	fmt.Println(Now().UnixNano())
	fmt.Println(Now().Weekday())
	fmt.Println(Now().WeekdayStr(1))
	fmt.Println(Str2Time("2018-04-23 23:11:23", "Y-m-d H:i:s").PFormat("Y-m-d H:i:s"))
}

```

```html

format 字符
日  --- --- 
d  月份中的第几天，有前导零的 2 位数字 01 到 31  
D  星期中的第几天，文本表示，3 个字母 Mon 到 Sun  
j  月份中的第几天，没有前导零 1 到 31  
l（“L”的小写字母） 星期几，完整的文本格式 Sunday 到 Saturday  
z  年份中的第几天 0 到 365  
星期  --- --- 
W  ISO-8601 格式年份中的第几周，每周从星期一开始 例如：42（当年的第 42 周） 
月  --- --- 
F  月份，完整的文本格式，例如 January 或者 March January 到 December  
m  数字表示的月份，有前导零 01 到 12  
M  三个字母缩写表示的月份 Jan 到 Dec  
n  数字表示的月份，没有前导零 1 到 12  
年  --- --- 
Y  4 位数字完整表示的年份 例如：1999 或 2003  
y  2 位数字表示的年份 例如：99 或 03  
时间  --- --- 
a  小写的上午和下午值 am 或 pm  
A  大写的上午和下午值 AM 或 PM  
g  小时，12 小时格式，没有前导零 1 到 12  
h  小时，12 小时格式，有前导零 01 到 12  
H  小时，24 小时格式，有前导零 00 到 23  
i  有前导零的分钟数 00 到 59> 
s  秒数，有前导零 00 到 59> 
u  毫秒，.000000 (比php多一个小数点)
时区  --- --- 
O  与格林威治时间相差的小时数 例如：+0200  
P  与格林威治时间（GMT）的差别，小时和分钟之间有冒号分隔 例如：+02:00  
T  本机所在的时区 例如：EST，MDT
完整的日期／时间  --- --- 
c  ISO 8601 格式的日期 2004-02-12T15:19:21+00:00 
r  RFC 822 格式的日期 例如：Thu, 21 Dec 2000 16:01:07 +0200  

```

