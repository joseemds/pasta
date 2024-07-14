package noodles

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func Routes(r chi.Router) {
	r.Post("/", createNoodle)
}

func createNoodle(w http.ResponseWriter, r *http.Request) {
	log := log.Default()

	defer r.Body.Close()

	validate := validator.New()
	req := new(CreateNoodleRequestBody)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error when parsing json %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, errors.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
