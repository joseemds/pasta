package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joseemds/pasta/internal/noodles"
	"go.uber.org/zap"
)


type App struct {
	Router *chi.Mux
	Logger *zap.Logger
}


func NewApp(logger *zap.Logger) *App {
	app := &App {
		Router: chi.NewRouter(),
		Logger: logger,
	}

	app.setupRoutes()

	return app
}

func (app *App) setupRoutes(){
	r := app.Router

  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))
	r.Route("/api/", func(r chi.Router) {
		r.Route("/noodles/", noodles.Routes)
	})

}
