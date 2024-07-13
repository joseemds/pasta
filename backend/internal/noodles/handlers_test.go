package noodles_test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
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
	}
	


	app := fiber.New()
	nods := app.Group("/api/noodles")
	noodles.RegisterNoodleGroup(nods)

	for _, tc := range testCases{
		t.Run(tc.description, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/noodles", bytes.NewReader(tc.payload))
			req.Header.Set("Content-Type", "application/json")


			res, _ := app.Test(req)

			assert.Equalf(t, tc.expectedCode, res.StatusCode, tc.description)
		})
	}

}

