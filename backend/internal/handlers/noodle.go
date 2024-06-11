package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/joseemds/pasta/internal/noodle"
)

func PostNoodles(w http.ResponseWriter, r *http.Request){
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	err := json.Marshal(noodles)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		l.Warn(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
