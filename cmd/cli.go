package cram

import "flag"

// variables used in flags
var (
	dbQuery string
	gui bool

	// this is not a cli argument
	resultChan chan []string
)

// cli args are intialized here
func InitCliArgs() {
	flag.StringVar(&dbQuery, "query", "", "query to use for querying the database")
	flag.BoolVar(&gui, "gui", true, "start the gui")
	flag.Parse()
}

//below are getters for cli args because they are private variables
func GetDbQuery() string {
	return dbQuery
}

func GetGui() bool {
	return gui
}

func GetResultChan() chan []string {
	return resultChan
}
