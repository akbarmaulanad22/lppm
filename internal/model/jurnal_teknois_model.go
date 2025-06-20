package model

type JurnalTeknoisResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateJurnalTeknoisRequest struct {
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required"`
}

type UpdateJurnalTeknoisRequest struct {
	ID      uint   `json:"-"`
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required"`
}

type DeleteJurnalTeknoisRequest struct {
	ID uint `json:"-" validate:"required"`
}
