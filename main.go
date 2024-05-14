package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thomaslievre/my-simple-bank/api"
	db "github.com/thomaslievre/my-simple-bank/db/sqlc"
)

const (
	dbSource   = "postgresql://bankroot:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddr = "0.0.0.0:8080"
)

func main() {
	config, err := pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}
	conn, err := pgxpool.NewWithConfig(context.Background(), config)

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddr)

	if err != nil {
		log.Fatal("cannot start server : ", err)
	}
}
