package main

import "flag"

// variables used in flags
var (
	databaseQuery string
	dbUsername string
	gui bool
)

// cli args are intialized here
func InitCliArgs() {
	flag.StringVar(&databaseQuery, "query", "", "query to use for querying the database")
	flag.StringVar(&dbUsername, "u", "", "the username for the database")
	flag.BoolVar(&gui, "gui", true, "start the gui")
	flag.Parse()
}

//below are getters for cli args because they are private variables
func GetDbQuery() string {
	return databaseQuery
}

func GetDbUsername() string {
	return dbUsername
}

func GetGui() bool {
	return gui
}
