package main

import "database/sql"

// return a channel to the connection
func connectToDb() {
	// would be what keeps the connection alive while the gui is running,
	// use channel to use connection, doesnt make sense for cli unless
	// planning to use daemon or background process
	if GetGui() {
		go func() {
			
		}()
	}
}

// Takes os or software and returns cve number, vulnerability score and time to fix
func QueryDbOS(os string) (string, float32, int){

}

// takes cve number and returns vulnerability score and time to fix
func QueryDbCve(cveNum float32) (float32, int){
	
}
