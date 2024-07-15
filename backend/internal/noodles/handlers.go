package noodles

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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

func (h NoodleHandler) Routes(r chi.Router) {
	r.Post("/", h.createNoodle)
}
func (h NoodleHandler) createNoodle(w http.ResponseWriter, r *http.Request) {
	log := h.Logger

	defer r.Body.Close()
	defer h.Logger.Sync()

	validate := validator.New()
	req := new(CreateNoodleRequestBody)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("Error when parsing json %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, errors.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := h.Service.createNoodles(req.Noodles); err != nil{
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		h.Logger.Errorf("Error creating noodle %w", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
