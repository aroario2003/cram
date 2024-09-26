package cram

import "flag"

// variables used in flags
var (
	dbName string
	dbQuery string
	dbUsername string
	gui bool
)

// cli args are intialized here
func InitCliArgs() {
	flag.StringVar(&dbQuery, "query", "", "query to use for querying the database")
	flag.StringVar(&dbUsername, "u", "", "the username for the database")
	flag.StringVar(&dbName, "n", "", "the name of the database to query")
	flag.BoolVar(&gui, "gui", true, "start the gui")
	flag.Parse()
}

//below are getters for cli args because they are private variables
func GetDbQuery() string {
	return dbQuery
}

func GetDbUsername() string {
	return dbUsername
}

func GetDbName() string {
	return dbName
}

func GetGui() bool {
	return gui
}
