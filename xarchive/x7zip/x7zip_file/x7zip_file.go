package x7zip_file

import (
	"github.com/Andyfoo/go-xutils/xencode"
	"github.com/Andyfoo/go-xutils/xfile"
)

func Init7zBin(path string) {
	var f_7z_exe = path + "/" + "7z.exe"
	var f_7z_dll = path + "/" + "7z.dll"
	if !xfile.FileIsExist(path) {
		xfile.MkdirAll(path)
	}
	if !xfile.FileIsExist(f_7z_exe) {
		xfile.WriteFile(f_7z_exe, xencode.Base64Decode([]byte(DATA_7z_exe)))
	}
	if !xfile.FileIsExist(f_7z_dll) {
		xfile.WriteFile(f_7z_dll, xencode.Base64Decode([]byte(DATA_7z_dll)))
	}
}
