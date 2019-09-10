package ping

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
	//"github.com/Andyfoo/go-xutils/xlog"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

var (
	icmp    ICMP
	timeout = 1500
	size    = 32
	num     = 4
)

func IsPingOK(desIp string) bool {

	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	//icmp头部填充
	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 1
	icmp.SequenceNum = 1

	//xlog.Infof("\n正在 ping %s 具有 %d 字节的数据:\n", desIp, size)

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp) // 以大端模式写入
	data := make([]byte, size)                    //
	buffer.Write(data)
	data = buffer.Bytes()

	for i := 0; i < num; i++ {
		icmp.SequenceNum = uint16(1)
		// 检验和设为0
		data[2] = byte(0)
		data[3] = byte(0)

		data[6] = byte(icmp.SequenceNum >> 8)
		data[7] = byte(icmp.SequenceNum)
		icmp.Checksum = CheckSum(data)
		data[2] = byte(icmp.Checksum >> 8)
		data[3] = byte(icmp.Checksum)

		// 开始时间
		t1 := time.Now()
		conn.SetDeadline(t1.Add(time.Duration(time.Duration(timeout) * time.Millisecond)))
		_, err := conn.Write(data)
		if err != nil {
			log.Fatal(err)
			return false
		}
		buf := make([]byte, 65535)
		_, err = conn.Read(buf)
		if err != nil {
			//xlog.Error("请求超时。", err)
			return false
		}
		//et := int(time.Since(t1) / 1000000)

		//xlog.Infof("来自 %s 的回复: 字节=%d 时间=%dms TTL=%d\n", desIp, len(buf[28:n]), et, buf[8])
		return true
	}
	return false
}

func CheckSum(data []byte) uint16 {
	var sum uint32
	var length = len(data)
	var index int
	for length > 1 { // 溢出部分直接去除
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length == 1 {
		sum += uint32(data[index])
	}
	// CheckSum的值是16位，计算是将高16位加低16位，得到的结果进行重复以该方式进行计算，直到高16位为0
	/*
		sum的最大情况是：ffffffff
		第一次高16位+低16位：ffff + ffff = 1fffe
		第二次高16位+低16位：0001 + fffe = ffff
		即推出一个结论，只要第一次高16位+低16位的结果，再进行之前的计算结果用到高16位+低16位，即可处理溢出情况
	*/
	sum = uint32(sum>>16) + uint32(sum)
	sum = uint32(sum>>16) + uint32(sum)
	return uint16(^sum)
}

func IsPortOK(ip string, port int) bool {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		//xlog.Error(err)
		return false
	}
	conn, err := net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		//	xlog.Error(err)
		return false
	} else {
		conn.Close()
		return true
	}

}
