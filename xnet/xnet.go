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

package xnet

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

//获取本机IP
func LocIP() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		//fmt.Println("net.Interfaces failed, err:", err.Error())
		return ""
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return ""
}

//获取当前机器互联网外网IP
func WanIP() string {
	//https://www.my-ip.io/api
	//https://api.my-ip.io/ip		IPv4, IPv6
	//https://api4.my-ip.io/ip		IPv4
	//https://api6.my-ip.io/ip		IPv6

	//https://ip-api.com/

	//https://seeip.org/
	//https://ip.seeip.org/   	IPv4 or IPv6
	//https://ip4.seeip.org/	IPv4
	//https://ip6.seeip.org		IPv6

	//https://ip.tool.lu/
	//http://ifconfig.me/ip
	//http://ipinfo.io/ip
	//https://ifconfig.co/ip
	resp, err := http.Get("https://ip4.seeip.org/")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	//s := buf.String()
	return string(content)
}

//判断是否是公网ip
func IsWanIP(ip string) bool {
	IPAddr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return false
	}
	IP := IPAddr.IP

	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

//ip地址string转int
func Ip2int(ipnr string) int64 {
	bits := strings.Split(ipnr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

//ip地址int转string
func Int2ip(ipnr int64) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String()
}

//判断ip地址区间
func IpBetween(fromIp string, toIp string, testIp string) bool {
	fromInt := Ip2int(fromIp)
	toInt := Ip2int(toIp)
	testInt := Ip2int(testIp)

	if fromInt == 0 || toInt == 0 || testInt == 0 {
		//fmt.Println("An ip did not convert to a 16 byte") // or return an error!?
		return false
	}

	if testInt >= fromInt && testInt <= toInt {
		return true
	}
	return false
}
