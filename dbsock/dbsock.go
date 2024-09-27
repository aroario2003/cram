// +build linux windows

package main

import (
	"flag"
	"log"
)

var (
	dbName string
	dbUsername string
)

// this takes cli args for connecting to the database
func InitCliArgs() {
	flag.StringVar(&dbName, "n", "", "the name of the database to connect to")
	flag.StringVar(&dbUsername, "u", "", "the username of the user to connect")
	flag.Parse()
}

func main() {
	InitCliArgs()
	if dbUsername == "" || dbName == "" {
		log.Fatalf("Cannot connect to database correctly, did not specify inforamtion, please use: dbsock -u <user> -n <db-name> to ensure correct functionality")
	}
	CreateSocket(dbName, dbUsername)
}
