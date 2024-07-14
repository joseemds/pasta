package noodles_test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/joseemds/pasta/internal/noodles"
	"github.com/stretchr/testify/assert"
)


func TestCreateNoodle(t *testing.T){
	testCases := []struct {
		description string
		payload []byte
		errorMessage string
		expectedCode int
		}{
		{
			description: "Create Noodles with EmptyList fails",
			payload:      []byte(`{"noodles": []}`),
			expectedCode: 422,
			errorMessage: "Expected NonEmpty list",
		},
		{
			description: "Create Single Noodle Succeed",
			payload:      []byte(`{"noodles": []}`),
			expectedCode: 201,
			errorMessage: "",
		},
	}
	



	for _, tc := range testCases{
		t.Run(tc.description, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/noodles", bytes.NewReader(tc.payload))
			req.Header.Set("Content-Type", "application/json")
			router := chi.NewRouter()
			router.Route("/api/noodles", noodles.Routes)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)


			assert.Equalf(t, tc.expectedCode, rr.Code, tc.description)
		})
	}

}

