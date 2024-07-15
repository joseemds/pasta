package pasta

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)


type PastaHandler struct {
	Conn *sql.DB
	Logger *zap.SugaredLogger
	Service *PastaService
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

func (h PastaHandler) createPasta(w http.ResponseWriter, r *http.Request){
	log := h.Logger

	defer r.Body.Close()
	defer h.Logger.Sync()
	validate := validator.New()
	reqBody := new(CreatePastaRequestBody)
	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		log.Errorf("Error when parsing json to create pasta %w", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(reqBody); err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, errors.Error(), http.StatusUnprocessableEntity)
		return
	}

if err := h.Service.CreatePasta(*reqBody); err != nil{
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		h.Logger.Errorf("Error creating noodle %w", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
