package main

import (
	"database/sql"
	"io"
	"log"
	"net"
	"fmt"
	"flag"

	cram "github.com/aroario2003/cram/cmd"
	_ "github.com/go-sql-driver/mysql"
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

// creates connection to the database
func connectToDb() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s@unix(/var/run/mysqld/mysqld.sock)/%s", dbUsername, dbName))
	if err != nil {
		log.Fatalf("Could not establish database connection: %v", err)
	}
	
	return db
}

// creates a unix domain socket to keep the database connection alive even after running the cli
func createSocketLinux() {
	// define the path to the socket for database connection
	socketPath := "/tmp/dbsock.sock"
	// create the socket and listen
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Could not create socket for database connection: %v", err)
	}
	// make sure the listener gets closed
	defer listener.Close()
	
	db := connectToDb()
	defer db.Close()

	log.Println("database socket created, waiting for queries...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Could not accept connection: %v", err)
			continue
		}
		
		// handle each connection to the socket.
		// this runs asynchronously
		go handleConnectionLinux(conn, db)

	}
}

// handles each connection to the unix doamin socket created above
func handleConnectionLinux(conn net.Conn, db *sql.DB) {
	defer conn.Close()
	
	// buffer for query to db, make it bigger?
	queryBuf := make([]byte, 1024)
	n, err := conn.Read(queryBuf)
	if err != nil && err != io.EOF {
		log.Printf("Failed to read query from connection: %v", err)
	}
	
	// get the query as a string
	query := string(queryBuf[:n])
	log.Printf("Executing Query: %s", query)
	// get rows from query
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to execute query: %s because %v", query, err)
	}
	// make sure that the table connection is closed
	defer rows.Close()
	
	// create a var for and print the resulting rows
	var result string
	for rows.Next() {
		if err := rows.Scan(&result); err != nil {
			log.Printf("Could not read row in result of query: %v", err)
		}
		conn.Write([]byte(fmt.Sprintf("%s", result)))
	}
}

// Creates a named pipe on windows to keep database connection
// alive even when cli isnt running
func createSocketWindows() {
	pipePath := "\\.\\pipe\\dbsockpipe"
	listener, err := net.Listen("npipe", pipePath)
	if err != nil {
		log.Fatalf("Failed to create named pipe: %v", err)
	}
	defer listener.Close()
	log.Printf("Named pipe created, listening...")
	
	db := connectToDb()
	defer db.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Could not accept connection to named pipe: %v", err)
		}
		
		// handle each connection to the socket.
		// this runs asynchronously
		go handleConnectionWindows(conn, db)
	}
}

// handles the connection to the named pipe on windows
func handleConnectionWindows(conn net.Conn, db *sql.DB) {
	defer conn.Close()
	
	queryBuf := make([]byte, 1024)
	n, err := conn.Read(queryBuf)
	if err != nil && err != io.EOF {
		log.Printf("Failed to read query from connection: %v", err)
	}

	// get the query as a string
	query := string(queryBuf[:n])
	log.Printf("Executing Query: %s", query)

	// get rows from query
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to execute query: %s because %v", query, err)
	}
	// make sure that the table connection is closed
	defer rows.Close()
	
	// create a var for and print the resulting rows
	var result string
	for rows.Next() {
		if err := rows.Scan(&result); err != nil {
			log.Printf("Could not read row in result of query: %v", err)
		}
		conn.Write([]byte(fmt.Sprintf("%s", result)))
	}
}

func main() {
	InitCliArgs()
	if cram.GetOs() == "linux" {
		createSocketLinux()
	} else if cram.GetOs() == "windows" {
		createSocketWindows()
	}
}
