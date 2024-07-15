package noodles

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type NoodleHandler struct {
	Logger *zap.SugaredLogger
	DBConn *sql.DB
	Service *NoodleService 
}

func NewHandler(logger *zap.SugaredLogger, conn *sql.DB) NoodleHandler{
	service := NewService(logger, conn)

	return NoodleHandler{
		Logger: logger,
		DBConn: conn,
		Service: &service,
	}
}

func (h NoodleHandler) Routes(r chi.Router) {}

