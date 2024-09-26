package dbsock

import (
	"database/sql"
	"io"
	"log"
	"net"
	"fmt"

	cram "github.com/aroario2003/cram/src"
	_ "github.com/go-sql-driver/mysql"
)

// creates a unix domain socket to keep the database connection alive even after running the cli
func createSocket() {
	// define the path to the socket for database connection
	socketPath := "/tmp/dbsock.sock"
	// create the socket and listen
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Could not create socket for database connection: %v", err)
	}
	// make sure the listener gets closed
	defer listener.Close()
	
	// connect to database
	dbUsername := cram.GetDbUsername()
	dbName := cram.GetDbName()
	db, err := sql.Open("mysql", fmt.Sprintf("%s@unix(/var/run/mysqld/mysqld.sock)/%s", dbUsername, dbName))
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
		
		// handle each connection to the socket
		go handleConnection(conn, db)

	}
}

// handles each connection to the unix doamin socket created above
func handleConnection(conn net.Conn, db *sql.DB) {
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

func main() {
	createSocket()
}
