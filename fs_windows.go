// +build windows

package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func Fsroot() string {
	return `C:\`
}

var (
	modkernel32 = windows.MustLoadDLL("kernel32.dll")

	GetDiskFreeSpaceEx = modkernel32.MustFindProc("GetDiskFreeSpaceExW")
)

func DiskAvail(path string) (uint64, error) {
	var size int64
	upath := windows.StringToUTF16Ptr(path)
	if r, _, err := GetDiskFreeSpaceEx.Call(uintptr(unsafe.Pointer(upath)), 0, 0, uintptr(unsafe.Pointer(&size))); r == 0 {
		return 0, err
	}
	return uint64(size), nil
}
