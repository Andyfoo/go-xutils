package ping

import (
	"fmt"
	"testing"
)

func TestIsPing(t *testing.T) {
	fmt.Println("IsPingOK", IsPingOK("192.168.1.1"))
	fmt.Println("IsPingOK", IsPingOK("192.168.1.2"))

}
func TestIsPort(t *testing.T) {
	fmt.Println("IsPortOK", IsPortOK("192.168.1.1", 80))
	fmt.Println("IsPortOK", IsPortOK("192.168.1.1", 443))

}
