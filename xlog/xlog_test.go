package xlog

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func xxx(file string) string {

	pos := strings.LastIndex(file, "/")
	if pos != -1 {
		pos1 := strings.LastIndex(file[:pos], "/src/")
		if pos1 != -1 {
			return file[pos1+5 : pos]
		}
	}
	type em struct{}
	file = reflect.TypeOf(em{}).PkgPath()
	pos = strings.Index(file, "/")
	if pos != -1 {
		return file[0:pos]
	}
	return "UNKNOWN_MOD"
}
func TestLog(t *testing.T) {
	fmt.Println(xxx("aaa.com/aaa/bbb"))
	SetOutputLevel(Ldebug)

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Warnf("Warn: foo\n")
	Warn("Warn: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")

	SetOutputLevel(Linfo)

	Debugf("Debug: foo\n")
	Debug("Debug: foo")

	Infof("Info: foo\n")
	Info("Info: foo")

	Warnf("Warn: foo\n")
	Warn("Warn: foo")

	Errorf("Error: foo\n")
	Error("Error: foo")
}
