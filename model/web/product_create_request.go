package web

type ProductCreateRequest struct {
	// Fields
	ID          string `json:"id" validate:"required,min=1"`
	Name        string `json:"name" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
	Image       string `json:"image"`
	CompanyID   int    `json:"company_id" validate:"required,min=1"`
}
