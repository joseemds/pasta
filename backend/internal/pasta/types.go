package pasta

import "github.com/joseemds/pasta/internal/noodles"

type PastaSchema struct {
	Noodles []noodles.NoodleSchema
}

type CreatePastaRequest struct {
	Noodles []PastaSchema `json:"noodles" validate:"min=1"`
}
