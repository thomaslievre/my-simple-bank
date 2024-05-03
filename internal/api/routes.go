package api

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello tout le monde titi\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

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

func (s *Server) RegisterRoutes() *http.ServeMux {
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)

	// r.Get("/", s.helloWorldHandler)

	// r.Get("/auth/{provider}/callback", s.getAuthProviderCallback)
	// r.Get("/logout/{provider}", s.logoutProvider)
	// r.Get("/auth/{provider}", s.beginAuthProviderCallback)
	// r.Get("/debug/providers", s.debugProvidersHandler)

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)

	return mux
}
