//go:build windows
// +build windows

package cram

import (
	"log"
	"github.com/Ne0nd0g/npipe"
)

func ConnectToDbSocket() *npipe.PipeConn {
	conn, err := npipe.Dial(`\\.\pipe\dbsockpipe`)
	if err != nil {
		log.Fatalf("Could not connect to the named pipe for dbsock: %v", err)
	}
	
	return conn
}

func GetDbConnection() *npipe.PipeConn {
	conn := ConnectToDbSocket()	
	return conn
}
