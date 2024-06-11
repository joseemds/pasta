package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joseemds/pasta/internal/noodle"
)

func TestPostNoodles(t *testing.T) {


	noodle := noodle.Noodle {
		Language: "python",
		Content: "println(\"Hello World\")",
		
	}

	body, err := json.Marshal(noodle)

	print(body)

	if err != nil {
		t.Error("unable to POST to /noodle")
	}

	req := httptest.NewRequest(http.MethodPost, "/noodle", nil)
	w := httptest.NewRecorder()
	handlers.PostNoodles(w, req)

	res := w.Result()


	if res.StatusCode !=  http.StatusCreated {
		t.Errorf("Expected 201 as a status code, received %d", res.StatusCode)
	}

	defer res.Body.Close()


}
