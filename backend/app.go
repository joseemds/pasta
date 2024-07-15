package main

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joseemds/pasta/internal/noodles"
	"github.com/joseemds/pasta/internal/pasta"
	"go.uber.org/zap"
)

type App struct {
	Router *chi.Mux
	Logger *zap.SugaredLogger
	DBConnection *sql.DB
}

func NewApp(logger *zap.SugaredLogger, conn *sql.DB) *App {
	app := &App{
		Router: chi.NewRouter(),
		Logger: logger,
		DBConnection: conn,
	}

	app.setupRoutes()

	return app
}

func (app *App) setupRoutes() {
	r := app.Router

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))
	r.Route("/api/", func(r chi.Router) {
		noodleHandler := noodles.NewHandler(app.Logger, app.DBConnection)
		pastaHandler := pasta.NewHandler(app.Logger, app.DBConnection)
		r.Route("/noodles", noodleHandler.Routes)
		r.Route("/pasta", pastaHandler.Routes)
	})

}
