package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/thomaslievre/my-simple-bank/db/sqlc"
	"github.com/thomaslievre/my-simple-bank/internal/token"
	"github.com/thomaslievre/my-simple-bank/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func ConnectToDB(dbSource string) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(dbSource)

	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)

	if err != nil {
		return nil, err
	}

	// defer dbPool.Close()

	return dbPool, nil
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.RegisterRoutes()

	return server, nil
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// func connectToDB() (*sql.DB, error) {
// 	const (
// 		host     = "localhost"
// 		port     = 5432
// 		user     = "root"
// 		password = "secret"
// 		dbname   = "simple_bank"
// 	)

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Successfully connected to database")
// 	return db, nil
// }

// func NewServer() *http.Server {
// 	// db, err := connectToDB()
// 	// if err != nil {
// 	// 	log.Fatalf("Could not connect to database: %v", err)
// 	// }

// 	// err = initTestTable(db)
// 	// if err != nil {
// 	// 	log.Fatalf("Could not initialize test table: %v", err)
// 	// }

// 	// Insert a test row
// 	// err = insertTestRow(db)
// 	// if err != nil {
// 	// 	log.Fatalf("Could not insert test row: %v", err)
// 	// }

// 	NewServer := &Server{
// 		port: port,
// 		db:   nil,
// 	}

// 	// Declare Server config
// 	server := &http.Server{
// 		Addr:         fmt.Sprintf(":%d", 8080),
// 		Handler:      NewServer.RegisterRoutes(),
// 		IdleTimeout:  time.Minute,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 	}

// 	return server
// }
