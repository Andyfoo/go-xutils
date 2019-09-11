package x7zip

import (
	"fmt"
	"testing"

	"github.com/Andyfoo/go-xutils/xlog"
)

func Test1(t *testing.T) {
	xlog.Info("test")
	fmt.Println("test")
	var extName = "rar"
	var data []byte
	var objPath = "E:/_tmp/test1/"
	filelist := UnRar(extName, data, objPath)
	for k, v := range filelist {
		fmt.Println(k, v)
	}

}
