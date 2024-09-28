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
)

// Takes os or software and returns cve number, vulnerability score and time to fix
func QueryDbOS() [][]interface{} {
	conn := GetDbConnection()
	defer conn.Close()

	query := fmt.Sprintf("select CVE_Number, Vulnerability_Score, Time_to_Fix from %s;", GetTableName())
	_, err := conn.Write([]byte(query))
	if err != nil {
		log.Printf("Could not send query over connection: %v", err)
	}
	
	resultChan := GetResultChan()
	results := <-resultChan
	return results
}

// takes cve number and returns vulnerability score and time to fix
func QueryDbCve(cveNum string) [][]interface{} {
	conn := GetDbConnection()
	defer conn.Close()

	query := fmt.Sprintf("select Vulnerability_Score, Time_to_Fix from %s where CVE_Number = %s;", GetTableName(), cveNum)
	_, err := conn.Write([]byte(query))
	if err != nil {
		log.Printf("Could not send query over connection: %v", err)
	}
	
	resultChan := GetResultChan()
	results := <-resultChan
	return results
}
