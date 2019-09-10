package xnet

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("LocIP", GetLocIP())
	fmt.Println("IsInternetIP", IsInternetIP("192.168.1.1"))
	fmt.Println("IsInternetIP", IsInternetIP("8.8.8.8"))

	fmt.Println("InternetIP", GetInternetIP())
}
func TestIP_int2str(t *testing.T) {
	fmt.Println("IP_str2int", IP_str2int("192.168.1.1"))
	fmt.Println("IP_int2str", IP_int2str(IP_str2int("192.168.1.1")))
}
