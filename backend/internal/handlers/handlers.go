package handlers

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)


func Logging(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		l.Info(
			"Request received",
			slog.String("path", r.RequestURI),
			slog.String("method", r.Method),
		)
		next.ServeHTTP(w, r)
	})
}

func NotFound(w http.ResponseWriter, r *http.Request){
	log.Println("Received req at " + r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}
