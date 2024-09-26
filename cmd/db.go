package cram

import "net"

// return the connection the database socket
func connectToDbSocket() net.Conn {

}

// Takes os or software and returns cve number, vulnerability score and time to fix
func QueryDbOS(os string) (string, float32, int){

}

// takes cve number and returns vulnerability score and time to fix
func QueryDbCve(cveNum float32) (float32, int){
	
}
