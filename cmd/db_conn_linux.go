//go:build linux
// +build linux

package cram

import (
	"net"
	"log"
)

// return the connection the database socket
func ConnectToDbSocket() net.Conn {
	conn, err := net.Dial("unix", "/tmp/dbsock.sock")
	if err != nil {
		log.Fatalf("Could not establish connection with database socket: %v", err)
	}
	return conn
}
