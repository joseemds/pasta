package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joseemds/pasta/internal/noodles"
	"go.uber.org/zap"
)

func main() {
	log := zap.Must(zap.NewDevelopment())

	r := chi.NewRouter()
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))
	r.Route("/api/", func(r chi.Router) {
		r.Route("/noodles/", noodles.Routes)
	})

	log.Info("Starting application at port 8080")
	http.ListenAndServe(":8080", r)
}
