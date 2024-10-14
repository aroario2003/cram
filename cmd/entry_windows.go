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
		totalVulnScore := GetTotalVulnerabilityScore(result, numRows)
		totalTimeToFix := GetTotalTimeToFix(result)
		fmt.Printf("%v\n", result)
		fmt.Printf("%d rows returned\n", numRows)
		fmt.Printf("%f total vulnerability score\n", totalVulnScore)
		fmt.Printf("%d total time to fix\n", totalTimeToFix)
	} else if GetCveNum() != "" { 
		conn := ConnectToDbSocket()
		result := QueryDbCve(conn, GetCveNum())
		numRows:= CountRowsReturned(result)
		totalVulnScore := GetTotalVulnerabilityScore(result, numRows)
		totalTimeToFix := GetTotalTimeToFix(result)
		fmt.Printf("%v\n", result)
		fmt.Printf("%d rows returned\n", numRows)
		fmt.Printf("%f total vulnerability score\n", totalVulnScore)
		fmt.Printf("%d total time to fix\n", totalTimeToFix)
	} else if len(GetOss()) != 0 {
		results := QueryDbMultiOs(GetOss())
		for i, result := range results {
			numRows := CountRowsReturned(result)
			totalVulnScore := GetTotalVulnerabilityScore(result, numRows)
			totalTimeToFix := GetTotalTimeToFix(result)
			fmt.Printf("===============query %d===============\n", i+1)
			fmt.Printf("%v\n", result)
			fmt.Printf("%d rows returned\n", numRows)
			fmt.Printf("%f total vulnerability score\n", totalVulnScore)
			fmt.Printf("%d total time to fix\n", totalTimeToFix)
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
			totalVulnScore := GetTotalVulnerabilityScore(result, numRows)
			totalTimeToFix := GetTotalTimeToFix(result)
			fmt.Printf("===============query %d===============\n", i+1)
			fmt.Printf("%v\n", result)
			fmt.Printf("%d rows returned\n", numRows)
			fmt.Printf("%f total vulnerability score\n", totalVulnScore)
			fmt.Printf("%d total time to fix\n", totalTimeToFix)
			if i != len(results)-1 {
				fmt.Printf("======================================\n")
			} else {
				fmt.Printf("======================================")
			}
		}
	}
}
