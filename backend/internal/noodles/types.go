package noodles

type NoodleSchema struct {
	Content   string `json:"content"`
	Filename  string `json:"filename"`
	Language  string `json:"language"`
}

type CreateNoodleRequestBody struct {
		Noodles []NoodleSchema `json:"noodles" validate:"min=1"`
}
