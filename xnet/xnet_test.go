package xnet

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("LocIP", LocIP())
	fmt.Println("IsWanIP", IsWanIP("192.168.1.1"))
	fmt.Println("IsWanIP", IsWanIP("8.8.8.8"))

	fmt.Println("WanIP", WanIP())
}
func TestIP_int2str(t *testing.T) {
	fmt.Println("Ip2int", Ip2int("192.168.1.1"))
	fmt.Println("Int2ip", Int2ip(Ip2int("192.168.1.1")))
	fmt.Println("IpBetween", IpBetween("192.168.1.1", "192.168.1.255", "192.168.2.2"))
}
