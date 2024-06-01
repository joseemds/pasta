package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.Use(logging)

	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "tome")
	})



	http.ListenAndServe("localhost:8090", r)
}
