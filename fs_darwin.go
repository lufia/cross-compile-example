// +build darwin

package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"golang.org/x/sys/unix"
)

func Fsroot() string {
	return "/"
}

func DiskAvail(path string) (uint64, error) {
	var s unix.Statfs_t
	if err := unix.Statfs(path, &s); err != nil {
		return 0, err
	}
	return s.Bavail * uint64(s.Bsize), nil
}
