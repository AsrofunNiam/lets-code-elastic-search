package domain

import (
	"encoding/json"
	"time"

	"github.com/AsrofunNiam/lets-code-elastic-search/model/web"
	"github.com/olivere/elastic/v7"
)

type Products []Product
type Product struct {
	ID          string `gorm:"primarykey;size:20"`
	CreatedByID uint   `gorm:""`
	CreatedAt   time.Time
	UpdatedByID uint `gorm:""`
	UpdatedAt   time.Time
	DeletedByID *uint     `gorm:""`
	DeletedAt   time.Time `gorm:"default:null"`

	Name        string `gorm:"type:varchar(100);not null"`
	Principal   string `gorm:"type:varchar(150);not null"`
	Description string `gorm:""`
	Image       string `gorm:"size:100"`
	CompanyID   int    `gorm:"primarykey"`
	Available   bool   `gorm:"default:true"`
}

func (product *Product) ToProductResponse() web.ProductResponse {
	return web.ProductResponse{
		ID:          product.ID,
		CreatedByID: product.CreatedByID,
		UpdatedByID: product.UpdatedByID,
		UpdatedAt:   product.UpdatedAt,
		Name:        product.Name,
		Description: product.Description,
		Image:       product.Image,
		CompanyID:   product.CompanyID,
	}
}

func (products Products) ToProductResponses() []web.ProductResponse {

	var productResponses []web.ProductResponse
	for _, dataVwProduct := range products {
		productResponses = append(productResponses, dataVwProduct.ToProductResponse())
	}
	return productResponses
}

func ToElasticProductResponses(hits []*elastic.SearchHit) []web.ProductResponse {
	var products []web.ProductResponse
	for _, hit := range hits {
		var product web.ProductResponse
		if err := json.Unmarshal(hit.Source, &product); err != nil {
			continue
		}
		products = append(products, product)
	}
	return products
}
