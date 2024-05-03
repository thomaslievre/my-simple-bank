package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDb *pgxpool.Pool

const (
	dbSource = "postgresql://bankroot:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// testDb, err = pgx.Connect(context.Background(), dbSource)

	poolConfig, err := pgxpool.ParseConfig(dbSource)
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
