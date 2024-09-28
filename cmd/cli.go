package cram

import "flag"

// variables used in flags
var (
	dbQuery string
	gui bool
	dbTable string
)

// cli args are intialized here
func InitCliArgs() {
	flag.StringVar(&dbQuery, "query", "", "query to use for querying the database")
	flag.StringVar(&dbTable, "t", "", "the name of the table to execute the query on")
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

func GetTableName() string {
	return dbTable
}

