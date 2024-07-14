package main

import (
	"net/http"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewDevelopment())

	app := NewApp(logger)

	app.Logger.Info("Starting application at port 8080")
	http.ListenAndServe(":8080", app.Router)
}
