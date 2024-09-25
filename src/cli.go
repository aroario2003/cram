package main

import "flag"

// variables used in flags
var (
	databaseQuery string
	gui bool
)

// cli args are intialized here
func InitCliArgs() {
	flag.StringVar(&databaseQuery, "query", "", "query to use for querying the database")
	flag.BoolVar(&gui, "gui", true, "start the gui")
	flag.Parse()
}
