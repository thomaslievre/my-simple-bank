package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thomaslievre/my-simple-bank/util"
)

var testQueries *Queries
var testDb *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	poolConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	testDb, err = pgxpool.NewWithConfig(context.Background(), poolConfig)

	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	// defer testDb.Close(context.Background())

	testQueries = New(testDb)

	os.Exit((m.Run()))
}
