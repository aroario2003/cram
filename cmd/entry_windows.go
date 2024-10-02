//go:build windows
// +build windows

package cram

import (
	"fmt"
)

func Entry() {
	InitCliArgs()

	if GetSoftware() != "" {
		conn := ConnectToDbSocket()
		result := QueryDbOS(conn, GetSoftware())
		numRows:= CountRowsReturned(result)
		fmt.Printf("%v", result)
		fmt.Printf("%d rows returned\n", numRows)
	} else if GetCveNum() != "" { 
		conn := ConnectToDbSocket()
		result := QueryDbCve(conn, GetCveNum())
		numRows:= CountRowsReturned(result)
		fmt.Printf("%v", result)
		fmt.Printf("%d rows returned\n", numRows)
	} else if len(GetOss()) != 0 {
		results := QueryDbMultiOs(GetOss())
		for i, result := range results {
			numRows := CountRowsReturned(result)
			fmt.Printf("===============query %d===============\n", i+1)
			fmt.Printf("%v", result)
			fmt.Printf("%d rows returned\n", numRows)
			if i != len(results)-1 {
				fmt.Printf("======================================\n")
			} else {
				fmt.Printf("======================================")
			}
		}
	} else if len(GetCveNums()) != 0 {
		results := QueryDbMultiCve(GetCveNums())
		for i, result := range results {
			numRows := CountRowsReturned(result)
			fmt.Printf("===============query %d===============\n", i+1)
			fmt.Printf("%v", result)
			fmt.Printf("%d rows returned\n", numRows)
			if i != len(results)-1 {
				fmt.Printf("======================================\n")
			} else {
				fmt.Printf("======================================")
			}
		}
	}
}
