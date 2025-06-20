package model

type PKMTCRResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

type CreatePKMTCRRequest struct {
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required"`
}

type UpdatePKMTCRRequest struct {
	ID      uint   `json:"-"`
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required"`
}

type DeletePKMTCRRequest struct {
	ID uint `json:"-" validate:"required"`
}
