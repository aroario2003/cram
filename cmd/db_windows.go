//go:build windows
// +build windows

// this code is duplicated but this is required because
// GetDbConnection returns different types based on ths 
// os and therefore requires this duplicated multi-platform code.
// Unduplicating this code will lead to platform specific errors.

package cram

import (
	"fmt"
	"log"
	"strings"
	"strconv"

	"github.com/Ne0nd0g/npipe"
)

func CountRowsReturned(result string) int {
	rows := strings.Split(result, "\n")
	return len(rows)
}

// get the total cvss score of all vulnerabilities returned by the query
func GetTotalVulnerabilityScore(result string, rowsCount int) float32 {
	var totalVulnScore float32
	var vulnScoreStr string
	resultsArr := strings.Split(result, "\n")

	for _, result := range resultsArr {
		resultArr := strings.Split(result, " ")
		
		if GetSoftware() != "" || len(GetOss()) != 0 {
			if len(resultArr) < 3 {
				log.Fatalf("The result does not have the correct amount of columns, your query paramaters may not exist or mysqld may not be started")
			}

			vulnScoreStr = resultArr[1]
		} else if GetCveNum() != "" || len(GetCveNums()) != 0 {
			if len(resultArr) < 2 {
				log.Fatalf("The result does not have the correct amount of columns, your query paramaters may not exist or mysqld may not be started")
			}

			vulnScoreStr = resultArr[0]
		} 

		vulnScore, err := strconv.ParseFloat(vulnScoreStr, 32)
		if err != nil {
			log.Printf("Could not convert string to float32 for total vulnerability score: %v", err)
		}

		totalVulnScore += float32(vulnScore) * float32(vulnScore)

	}

	rawVulnScore := totalVulnScore / float32(rowsCount)
	vulnScore := 100 - rawVulnScore

	return vulnScore
}

// gets the total time to fix of all vulnerabilities returned by the query
func GetTotalTimeToFix(result string) uint16 {
	var totalTimeToFix uint16
	var ttfStr string
	resultsArr := strings.Split(result, "\n")

	for _, result := range resultsArr {
		resultArr := strings.Split(result, " ")

		if GetSoftware() != "" || len(GetOss()) != 0 {
			if len(resultArr) < 3 {
				log.Fatalf("The result does not have the correct amount of columns, your query paramaters may not exist or mysqld may not be started")
			}

			ttfStr = resultArr[2]
		} else if GetCveNum() != "" || len(GetCveNums()) != 0 {
			if len(resultArr) < 2 {
				log.Fatalf("The result does not have the correct amount of columns, your query paramaters may not exist or mysqld may not be started")
			}

			ttfStr = resultArr[1]
		}

		ttf, err := strconv.ParseUint(ttfStr, 10, 8)
		if err != nil {
			log.Printf("Could not convert string to uint8 for total time to fix: %v", err)
		}

		totalTimeToFix += uint16(ttf)
	}

	minutes := (totalTimeToFix % 3600) / 60
	hours := totalTimeToFix / 3600

	if hours == 0 && minutes == 0 { 
		totalTimeToFix = 0 
	} else if hours == 0 && minutes == 1 { 
		totalTimeToFix = 1 
	} else if hours == 0 && minutes <= 5 {
		totalTimeToFix = 2
	} else if hours == 0 && minutes <= 15 {
		totalTimeToFix = 3 
	} else if hours == 0 && minutes <= 30 {
		totalTimeToFix = 4 
	} else if hours == 1 && minutes == 0 {
		totalTimeToFix = 5
	} else if hours == 2 && minutes == 0 {
		totalTimeToFix = 6
	} else if hours <= 4 {
		totalTimeToFix = 7
	} else if hours <= 8 {
		totalTimeToFix = 8
	} else if hours <= 12 {
		totalTimeToFix = 9
	} else if hours <= 24 {
		totalTimeToFix = 10
	}

	return totalTimeToFix
}

// Takes os or software and returns cve number, vulnerability score and time to fix
func QueryDbOS(conn *npipe.PipeConn, os string) string {
	defer conn.Close()

	query := fmt.Sprintf("select CVE_Number, Vulnerability_Score, Time_to_Fix from %s where Software like '%%%s%%' and Solved = 0", GetTableName(), os)
	_, err := conn.Write([]byte(query))
	if err != nil {
		log.Printf("Could not send query over connection: %v", err)
	}
	
	resultBuf := make([]byte, 10000000)
	n, err :=  conn.Read(resultBuf)
	resStr := string(resultBuf[:n])
	if err != nil {
		log.Printf("Could not read result from database socket: %v", err)
	}
	return resStr
}

// takes cve number and returns vulnerability score and time to fix
func QueryDbCve(conn *npipe.PipeConn, cveNum string) string {
	defer conn.Close()

	query := fmt.Sprintf("select Vulnerability_Score, Time_to_Fix from %s where CVE_Number = '%s' and Solved = 0", GetTableName(), cveNum)
	_, err := conn.Write([]byte(query))
	if err != nil {
		log.Printf("Could not send query over connection: %v", err)
	}
	
	resultBuf := make([]byte, 10000000)
	n, err :=  conn.Read(resultBuf)
	resStr := string(resultBuf[:n])
	if err != nil {
		log.Printf("Could not read result from database socket: %v", err)
	}
	return resStr
}

func QueryDbMultiOs(oss []string) []string {
	var results []string
	
	for _, os := range oss {
		conn := ConnectToDbSocket()
		defer conn.Close()

		query := fmt.Sprintf("select CVE_Number, Vulnerability_Score, Time_to_Fix from %s where Software like '%%%s%%' and Solved = 0", GetTableName(), os)

		_, err := conn.Write([]byte(query))
		if err != nil {
			log.Printf("Could not send query over connection: %v", err)
		}

		resultBuf := make([]byte, 10000000)
		n, err :=  conn.Read(resultBuf)
		if err != nil {
			log.Printf("Could not read result from database socket: %v", err)
		}

		resStr := string(resultBuf[:n])
		results = append(results, resStr)
	}
	return results
}

func QueryDbMultiCve(cves []string) []string {
	var results []string
	
	for _, cve := range cves {
		conn := ConnectToDbSocket()
		defer conn.Close()

		query := fmt.Sprintf("select Vulnerability_Score, Time_to_Fix from %s where CVE_Number = '%s' and Solved = 0", GetTableName(), cve)

		_, err := conn.Write([]byte(query))
		if err != nil {
			log.Printf("Could not send query over connection: %v", err)
		}

		resultBuf := make([]byte, 10000000)
		n, err :=  conn.Read(resultBuf)
		if err != nil {
			log.Printf("Could not read result from database socket: %v", err)
		}

		resStr := string(resultBuf[:n])
		results = append(results, resStr)
	}
	return results
}
