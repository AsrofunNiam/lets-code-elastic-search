package repository

import (
	"github.com/AsrofunNiam/lets-code-elastic-search/helper"
	"github.com/AsrofunNiam/lets-code-elastic-search/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) FindAll(db *gorm.DB) domain.Products {
	products := domain.Products{}
	tx := db.Model(&domain.Product{})
	err := tx.Find(&products).Error
	helper.PanicIfError(err)

	return products
}
func (repository *ProductRepositoryImpl) Create(db *gorm.DB, product *domain.Product) *domain.Product {
	err := db.Create(&product).Error
	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) Delete(db *gorm.DB, id *string) {
	product := &domain.Product{}
	// Deleting the leave from the database.
	err := db.Unscoped().Delete(product, id).Error
	helper.PanicIfError(err)
}
