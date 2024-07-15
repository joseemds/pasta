package pasta

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)


type PastaHandler struct {
	Conn *sql.DB
	Logger *zap.SugaredLogger
}


func NewHandler(logger *zap.SugaredLogger, conn *sql.DB) PastaHandler {
	return PastaHandler{
		Conn: conn,
		Logger: logger,
	}
}


func (h PastaHandler) Routes(r chi.Router){
	r.Post("/", h.createPasta)
}

func (h PastaHandler) createPasta(w http.ResponseWriter, r *http.Request){}
