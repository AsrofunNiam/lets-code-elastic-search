package web

import "time"

type ProductResponse struct {
	// Required Fields
	ID          string    `json:"id"`
	CreatedByID uint      `json:"created_by_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedByID uint      `json:"updated_by_id"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Fields
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CompanyID   int    `json:"company_id"`
	Available   bool   `json:"available"`
}
