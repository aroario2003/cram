//go:build linux || windows
// +build linux windows

package cram

import (
	"fmt"
)

func Entry() {
	InitCliArgs()
	conn := ConnectToDbSocket()
	result2 := QueryDbCve(conn, "CVE-2016-8024")
	fmt.Printf("%v", result2)
}
