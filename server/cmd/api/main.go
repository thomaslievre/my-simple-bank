package main

import (
	"log"

	db "github.com/thomaslievre/my-simple-bank/db/sqlc"
	server "github.com/thomaslievre/my-simple-bank/internal/api"

	"github.com/thomaslievre/my-simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbpool, err := server.ConnectToDB(config.DBSource)

	if err != nil {
		log.Fatal("Error during db connection: ", err)
	}

	store := db.NewStore(dbpool)
	server, err := server.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server : ", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server : ", err)
	}
}
