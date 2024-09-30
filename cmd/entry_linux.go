//go:build linux
// +build linux

package cram

import (
	"fmt"
)

func Entry() {
	InitCliArgs()
	conn := ConnectToDbSocket()

	if GetSoftware() != "" {
		result1 := QueryDbOS(conn, GetSoftware())
		fmt.Printf("%v", result1)
	}
	if GetCveNum() != "" { 
		result2 := QueryDbCve(conn, GetCveNum())
		fmt.Printf("%v", result2)
	}
}
