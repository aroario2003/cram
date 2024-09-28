//go:build windows
// +build windows

package main

import (
	"database/sql"
	"io"
	"log"
	"net"
	"fmt"

	"github.com/Ne0nd0g/npipe"
	cram "github.com/aroario2003/cram/cmd"
)

func createSocketListenerWindows(pipePath string) *npipe.PipeListener {
	listener, err := npipe.Listen(pipePath)
	if err != nil {
		log.Fatalf("Could not create named pipe: %v", err)
	}

	return listener
}

// Creates a named pipe on windows to keep database connection
// alive even when cli isnt running
func CreateSocket(dbName string, dbUsername string, dbPassword string) {
	pipePath := `\\.\pipe\dbsockpipe`

	listener := createSocketListenerWindows(pipePath)
	defer listener.Close()

	log.Printf("Named pipe created, listening...")
	
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@unix(/var/run/mysqld/mysqld.sock)/%s", dbUsername, dbPassword, dbName))
	if err != nil {
		log.Fatalf("Could not establish database connection: %v", err)
	}
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

