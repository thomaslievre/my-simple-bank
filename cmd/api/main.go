package main

import (
	server "github.com/thomaslievre/my-go-api/internal/api"
)

func main() {
	server := server.NewServer()

	// defer func() {
	// 	if err := server.Close(); err != nil {
	// 		log.Printf("Error closing the server: %v", err)
	// 	}
	// }()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}

}
