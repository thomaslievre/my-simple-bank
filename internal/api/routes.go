package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func messageHandler(message string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte(message))
// 	})
// }

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareOne again")
	})
}

func (s *Server) RegisterRoutes() {
	router := gin.Default()
	router.POST("/users", s.createUser)
	router.POST("/login", s.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))

	// authenticated routes
	authRoutes.POST("/accounts", s.createAccount)
	authRoutes.GET("/accounts/:id", s.getAccount)
	authRoutes.GET("/accounts", s.listAccount)
	authRoutes.PATCH("/accounts", s.updateAccount)
	authRoutes.DELETE("/accounts/:id", s.deleteAccount)
	authRoutes.POST("/transfers/create", s.createTransfer)

	s.router = router
}
