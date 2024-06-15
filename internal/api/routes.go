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

	// routes
	router.POST("/accounts", s.createAccount)
	router.GET("/accounts/:id", s.getAccount)
	router.GET("/accounts", s.listAccount)
	router.PATCH("/accounts", s.updateAccount)
	router.DELETE("/accounts/:id", s.deleteAccount)

	s.router = router
}
