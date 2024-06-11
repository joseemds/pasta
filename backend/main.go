package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joseemds/pasta/internal/handlers"
)


func logging(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received " + r.Method + " at " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func notFound(w http.ResponseWriter, r *http.Request){
	log.Println("Received req at " + r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	mux := http.NewServeMux();
	mux.HandleFunc("/health_check", handlers.HealthCheck)
	mux.HandleFunc("/noodle", handlers.PostNoodles)

	handler := handlers.Logging(mux)

	fmt.Println("Server started at port: 8090")
	log.Fatal(http.ListenAndServe("localhost:8090", handler))
}
