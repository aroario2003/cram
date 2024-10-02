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
	"strings"
)

func CountRowsReturned(result string) int {
	rows := strings.Split(result, "\n")
	return len(rows)
}

// Takes os or software and returns cve number, vulnerability score and time to fix
func QueryDbOS(conn net.Conn, os string) string {
	defer conn.Close()

	query := fmt.Sprintf("select CVE_Number, Vulnerability_Score, Time_to_Fix from %s where Software like '%%%s%%'", GetTableName(), os)

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

	query := fmt.Sprintf("select Vulnerability_Score, Time_to_Fix from %s where CVE_Number = '%s'", GetTableName(), cveNum)

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

		query := fmt.Sprintf("select CVE_Number, Vulnerability_Score, Time_to_Fix from %s where Software like '%%%s%%';", GetTableName(), os)

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

		query := fmt.Sprintf("select Vulnerability_Score, Time_to_Fix from %s where CVE_Number = '%s'", GetTableName(), cve)

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
