//go:build linux
// +build linux

package main

import (
	"database/sql"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

func createSocketListenerLinux(socketPath string) net.Listener {
	// create the socket and listen
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Could not create socket for database connection: %v", err)
	}
	return listener
}


// creates a unix domain socket to keep the database connection alive even after running the cli
func CreateSocket(dbName string, dbUsername string, dbPassword string) {
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
	
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@unix(/var/run/mysqld/mysqld.sock)/%s", dbUsername, dbPassword, dbName))
	if err != nil {
		log.Fatalf("Could not establish database connection: %v", err)
	}
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
					log.Printf("Could not write to database socket: %v", err)
				}
			} else {
				result = result + fmt.Sprintf("%v ", col)
			} 		
		}
		result += "\n"
	}

	_, err = conn.Write([]byte(result))
	if err != nil {
		log.Printf("Could not write results to database socket: %v", err)
	}
}

