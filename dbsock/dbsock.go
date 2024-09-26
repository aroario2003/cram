package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

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

func createSocketListenerLinux(socketPath string) net.Listener {
	// create the socket and listen
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Could not create socket for database connection: %v", err)
	}
	return listener
}

func createSocketListenerWindows(socketPath string) net.Listener {
	// create the socket and listen
	listener, err := net.Listen("npipe", socketPath)
	if err != nil {
		log.Fatalf("Could not create socket for database connection: %v", err)
	}
	return listener
}

// creates a unix domain socket to keep the database connection alive even after running the cli
func createSocketLinux() {
	// define the path to the socket for database connection
	socketPath := "/tmp/dbsock.sock"
	var listener net.Listener
	// check if the socket already exists
	if _, err := os.Stat(socketPath); errors.Is(err, os.ErrNotExist) {
		// if not create it
		listener = createSocketListenerLinux(socketPath)
		defer listener.Close()
	} else {
		// otherwise remove it and recreate it
		os.Remove(socketPath)
		listener = createSocketListenerLinux(socketPath)
		defer listener.Close()
	}
	
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
	var row string
	var result []string
	resultChan := cram.GetResultChan()
	for rows.Next() {
		if err := rows.Scan(&row); err != nil {
			log.Printf("Could not read row in result of query: %v", err)
		}
		result = append(result, row)
	}
	// send the result over the channel
	resultChan <- result
}

// Creates a named pipe on windows to keep database connection
// alive even when cli isnt running
func createSocketWindows() {
	pipePath := "\\.\\pipe\\dbsockpipe"
	var listener net.Listener
	// check if the socket already exists
	if _, err := os.Stat(pipePath); errors.Is(err, os.ErrNotExist) {
		// if not create it
		listener = createSocketListenerWindows(pipePath)
		defer listener.Close()
	} else {
		// otherwise remove it and recreate it
		os.Remove(pipePath)
		listener = createSocketListenerLinux(pipePath)
		defer listener.Close()
	}
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
	var result []string
	var row string
	resultChan := cram.GetResultChan()
	for rows.Next() {
		if err := rows.Scan(&result); err != nil {
			log.Printf("Could not read row in result of query: %v", err)
		}
		result = append(result, row)
	}
	resultChan <- result
}

func main() {
	InitCliArgs()
	if cram.GetOs() == "linux" {
		if dbUsername == "" || dbName == "" {
			log.Fatalf("Cannot connect to database correctly, did not specify inforamtion, please use: dbsock -u <user> -n <db-name> to ensure correct functionality")
		}
		createSocketLinux()
	} else if cram.GetOs() == "windows" {
		if dbUsername == "" || dbName == "" {
			log.Fatalf("Cannot connect to database correctly, did not specify inforamtion, please use: dbsock -u <user> -n <db-name> to ensure correct functionality")
		}
		createSocketWindows()
	}
}
