package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

var port = 8080

type Server struct {
	port int
	db   *sql.DB
}

func connectToDB() (*sql.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "root"
		password = "secret"
		dbname   = "simple_bank"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}

func NewServer() *http.Server {
	// db, err := connectToDB()
	// if err != nil {
	// 	log.Fatalf("Could not connect to database: %v", err)
	// }

	// err = initTestTable(db)
	// if err != nil {
	// 	log.Fatalf("Could not initialize test table: %v", err)
	// }

	// Insert a test row
	// err = insertTestRow(db)
	// if err != nil {
	// 	log.Fatalf("Could not insert test row: %v", err)
	// }

	NewServer := &Server{
		port: port,
		db:   nil,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", 8080),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
