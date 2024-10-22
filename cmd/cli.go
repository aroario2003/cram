package cram

import (
	"flag"
	"fmt"
	"os"
)

type argsArray []string

// variables used in flags
var (
	dbQuery string
	gui bool
	dbTable string
	software string
	cveNum string
	softwares argsArray
	cveNums argsArray
	markAsSolved string
)

// the two functions below are implementations of the interface
// of the flag package
func (a *argsArray) String() string {
    return fmt.Sprintf("%v", *a)
}

func (a *argsArray) Set(value string) error {
	*a = append(*a, value)
	return nil
}

// cli args are intialized here
func InitCliArgs() {
	flag.StringVar(&dbTable, "t", "", "the name of the table to execute the query on")
	flag.StringVar(&software, "s", "", "use the os/software query on the software/os specified")
	flag.StringVar(&cveNum, "c", "", "use the cve number query on the cve number specified")
	flag.Var(&softwares, "S", "multiple os's to query in a list")
	flag.Var(&cveNums, "C", "multiple cve's to query in a list")
	flag.StringVar(&markAsSolved, "m", "", "mark a cve as solved")
	flag.BoolVar(&gui, "gui", false, "start the gui")
	flag.Parse()

	parseSorCFlags()
}

// if the command contains one of -S or -C 
// then this is what does the work to collect the arguments
func parseSorCFlags() {
	var cves argsArray = nil
	var oss argsArray = nil

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-S" && i+1 < len(os.Args) {
			for j := i+1; j < len(os.Args); j++ {
				if os.Args[j][0] == '-' {
					break
				}
				oss = append(oss, os.Args[j])
			}
		} else if os.Args[i] == "-C" && i+1 < len(os.Args) {
			for j := i+1; j < len(os.Args); j++ {
				if os.Args[j][0] == '-' {
					break
				}
				cves = append(cves, os.Args[j])
			}
		}
	}

	if oss != nil {
		softwares = oss
	} else if cves != nil {
		cveNums = cves
	}
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

func GetSoftware() string {
	return software
}

func GetCveNum() string {
	return cveNum
}

func GetOss() argsArray {
	return softwares 
}

func GetCveNums() argsArray {
	return cveNums
}

func GetMarkAsSolved() string {
	return markAsSolved
}

// Below are setters for certain variables that are 
// command line arguments, this is the easiest way to 
// get the gui to work without too much refactoring
func SetSoftware(os string) {
	software = os
}

func SetCveNum(cve string) {
	cveNum = cve
}

func SetSoftwares(oss []string) {
	softwares = oss
}

func SetCveNums(cves []string) {
	cveNums = cves
}
