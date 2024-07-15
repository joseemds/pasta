package pasta

import "github.com/joseemds/pasta/internal/noodles"

type PastaSchema struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

type CreatePastaRequestBody struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Noodles []noodles.NoodleSchema `json:"noodles" validate:"min=1"`
}
