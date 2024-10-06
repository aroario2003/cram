# CRAM Cli 

## Introduction 

CRAM CLI streamlines the process of querying security vulnerabilities, offering a multi-platform solution for both Linux and Windows users. With a focus on efficiency, the tool maintains persistent database connections through Unix domain sockets or Windows named pipes. Users can retrieve CVE data, vulnerability scores, and time-to-fix details, making CRAM CLI an essential tool for security assessments in diverse environments.

# Security Notice

We are aware that when starting dbsock you must enter the password for the database in plain text and that could be a problem for a security critical environment.

# Software Versions Supported

- Cisco Firepower 4125 – Firepower Threat Defense (FTD) Software Version: 6.6.7
- Cisco Meraki MS425-32 Layer 3 Switch – Firmware: 2014-09-23
- Cisco Catalyst 2960-X Layer 2 Gigabit Ethernet Network Switch – IOS Version: 15.2(1)E
- RedHat Enterprise Linux (RHEL):  Version 5.0 Version 6.0 Version 7.1
- McAfee VirusScan Enterprise – Version: 2.0
- Tenable Nessus Vulnerability Scanner – Version: 8.10.0
- Splunk Security Information and Event Manager (SIEM) – Version: 8.6
- Microsoft Windows Server 2008 – Service Pack 2
- Apache OpenOffice (Open Source) – Version: 4.1.1.4

# Time To Fix Scale

The following is the scale for the total time to fix:

- 0: seconds to fix 
- 1: 1 minute to fix 
- 2: 5 minutes to fix 
- 3: 15 minutes to fix 
- 4: 30 minutes to fix
- 5: hour to fix 
- 6: 2 hours to fix 
- 7: 4 hours to fix 
- 8: 8 hours to fix
- 9: 12 hours to fix
- 10: full day to fix

## Multi-platform building

Go has a robust build system and therefore certain features are used to make this application multi-platform. To clarify, multi-platform means that this will not run natively on all platforms without certain modifications to the build system and codebase. You do not have to make these modifications as they are already made for you. However, it is worth knowing what they are.

## Go build tags and comment compiler directives

Build tags in go tell the compiler which version of your application you want to build, a go build tag takes the following form:

```
// +build linux
```

This tells the compiler that whenever it gets a tag linux, it should build files with that comment on top of them. You can also specify multiple tags in these comments.

There are also go compiler directives in comments, this tells the compiler which platform to compile for based on the current go runtime which boils down to the operating system. These take the form:

```
//go:build linux
```

This tells the compiler that it should build this file specifically when the runtime is linux, the same could be said for windows if it was windows in that comment.

### Code Duplication Side Effect

Because of the multi-platform nature of this application, their is quite a bit of code duplication in multiple functions throughout. However, this is absolutely necessary because functions will return different types depending on the operating system and therefore must be duplicated for the platforms supported.

## Installing Go

Refer to [this](https://go.dev/doc/install) page for instructions.

## What is dbsock?

In order to increase efficiency we create a program that creates a unix domain socket or named pipe on windows that maintains the database conection even while the cli isnt running. Dbsock is the program that creates that socket and connects it to the database.

## Building dbsock

### On Linux

In order to build dbsock on linux, from the root of the project enter these commands:

```shell
$ cd dbsock
$ go build -o dbsock
```

### On Windows

In order to build dbsock on windows, from the root of the project enter these commands:

```shell
$ cd dbsock
$ go build -o dbsock.exe
```

## Building the main program

### On Linux

From the root of the project do

```shell
$ go build -o main
```

### On Windows

From the root of the project do

```shell
$ go build -o main.exe
```

## Error When Building on Windows

When building on windows an error may occur that says to include the flag `-buildvcs=false`. If this happens just include the flag:

```shell
$ go build -o <program-name> -buildvcs=false
```

The `program-name` maybe either `dbsock` or `main`

You will most likely only have to do this the first time you build the programs.

## Using Dbsock

In order to start dbsock, to maintain a constant database connection, you must give the binary three command line arguments, first is the username of the user for the database software. Second is the name of the database that you want to use. Finally is the password of the user that you are logging into database software with. The full command resembles the following:

```shell
$ ./dbsock -u <username> -n <db-name> -p <password>
```

Once dbsock is started you should recieve a log message resembling the following on linux:

```
mm/dd/yyyy hh:mm:ss database socket created, waiting for queries...
```

Or on windows:

```
mm/dd/yyyy hh:mm:ss named pipe created, waiting for queries...
```

## Using The Cli

In order to execute queries from the cli you must give it the table name of the table to query and a flag to say which function to use to query. The command resembles the following:

```shell
$ ./main -t <table-name> -? ...
```

The `?` should be replaced with one of the following characters:

| Flag | Argument(s)
|------|-------------------------------------------------------|
| s    | Takes one argument, the name of the software          |
| c    | Takes one argument, the name of the cve               |
| S    | Takes variable arguments, the names of the softwares  |
| C    | Takes variable arguments, the names of the cves       |

The `...` represents the arguments to that flag 

## Credits

- Alfrickr - Created the database and worked on developing GUI
- Bluelightspirit - Contributed ideas for CLI and helped with testing on windows
