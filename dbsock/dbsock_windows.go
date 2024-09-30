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
		// get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Could not get column names of table: %v", err)
	}
	
	
	// create a var for and print the resulting rows
	var results [][]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}
	
		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal(err)
		}
		
		results = append(results, values)
	}

	var result string
	// Print the results (each row)
	for _, row := range results {
		for i, col := range row {
			if b, ok := col.([]byte); ok {
				if i+1 % len(columns) == 0 {
					result = result + fmt.Sprintf("%s", string(b))
				} else {
					result = result + fmt.Sprintf("%s ", string(b))
				}
				if err != nil {
					log.Printf("Could not write to database named pipe: %v", err)
				}
			} else {
				result = result + fmt.Sprintf("%v ", col)
			} 		
		}
		result += "\n"
	}

	_, err = conn.Write([]byte(result))
	if err != nil {
		log.Printf("Could not write results to database named pipe: %v", err)
	}
}

