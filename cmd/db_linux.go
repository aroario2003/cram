//go:build linux
// +build linux

// this code is duplicated but this is required because
// GetDbConnection returns different types based on ths
// os and therefore requires this duplicated multi-platform code
// Unduplicating this code will lead to platform specific errors.

package cram

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

// get the amount of rows returned by the query
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

		// weighting with linear scaling
		totalVulnScore += float32(vulnScore) * float32(vulnScore)
	}
	
	_, multiplier := GetTotalTimeToFix(result)
	rawVulnScore := (totalVulnScore / float32(rowsCount)) * multiplier
	if rawVulnScore > 100 {
		rawVulnScore = 100
	}
	vulnScore := 100 - rawVulnScore

	return vulnScore
}

// gets the total time to fix of all vulnerabilities returned by the query
func GetTotalTimeToFix(result string) (uint16, float32) {
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
	
	multiplier := float32(1.0)
	if totalTimeToFix > 80 { 
		multiplier = 1.3
	} else if totalTimeToFix <= 80 && totalTimeToFix > 60 {
		multiplier = 1.2
	} else if totalTimeToFix >= 50 && totalTimeToFix < 60 {
		multiplier = 1.1 
	} else if totalTimeToFix >= 40 && totalTimeToFix < 50 {
		multiplier = 1.0
	} else if totalTimeToFix >= 30 && totalTimeToFix < 40 {
		multiplier = 0.95
	} else if totalTimeToFix >= 20 && totalTimeToFix < 30 {
		multiplier = 0.9
	} else if totalTimeToFix < 20 {
		multiplier = 0.85
	}

	return totalTimeToFix, multiplier
}

// Takes os or software and returns cve number, vulnerability score and time to fix
func QueryDbOS(conn net.Conn, os string) string {
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
	return resStr
}

// takes cve number and returns vulnerability score and time to fix
func QueryDbCve(conn net.Conn, cveNum string) string {
	defer conn.Close()

	query := fmt.Sprintf("select Vulnerability_Score, Time_to_Fix from %s where CVE_Number = '%s' and Solved = 0", GetTableName(), cveNum)

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

func MarkAsSolved(cveName string) {
	conn := ConnectToDbSocket()
	defer conn.Close()
	
	query := fmt.Sprintf("update %s set Solved = 1 where CVE_Number = '%s'", GetTableName(), cveName)
	_, err := conn.Write([]byte(query))
	if err != nil {
		log.Printf("Could not send query over connection: %v", err)
	}
}
