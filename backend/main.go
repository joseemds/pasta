package main

import (
	"database/sql"
	"net/http"
	"os"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewDevelopment()).Sugar()
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil{
		panic(err)
	}

	app := NewApp(logger, db)
	app.Logger.Info("Starting application at port 8080")
	http.ListenAndServe(":8080", app.Router)
}
