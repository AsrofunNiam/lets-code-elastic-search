package repository

import (
	"github.com/AsrofunNiam/lets-code-elastic-search/model/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(db *gorm.DB) domain.Products
	Create(db *gorm.DB, product *domain.Product) *domain.Product
	Delete(db *gorm.DB, id *string)
}
