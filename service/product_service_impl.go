package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AsrofunNiam/lets-code-elastic-search/helper"
	"github.com/AsrofunNiam/lets-code-elastic-search/model/domain"
	"github.com/AsrofunNiam/lets-code-elastic-search/model/web"
	"github.com/AsrofunNiam/lets-code-elastic-search/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
	Validate          *validator.Validate
	ElasticClient     *elastic.Client
}

func NewProductService(
	product repository.ProductRepository,
	db *gorm.DB,
	validate *validator.Validate,
	elasticClient *elastic.Client,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository: product,
		DB:                db,
		Validate:          validate,
		ElasticClient:     elasticClient,
	}
}

func (service *ProductServiceImpl) Create(request *web.ProductCreateRequest, c *gin.Context) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	product := &domain.Product{
		// Required Fields
		ID: request.ID,

		// Fields
		Name:        request.Name,
		Description: request.Description,
		Image:       request.Image,
		CompanyID:   request.CompanyID,
	}
	product = service.ProductRepository.Create(tx, product)
	return product.ToProductResponse()
}

func (service *ProductServiceImpl) Delete(id *string, c *gin.Context) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	service.ProductRepository.Delete(tx, id)
}

func (service *ProductServiceImpl) Sync(c *gin.Context) {
	ctx := context.Background()
	bulkRequest := service.ElasticClient.Bulk()
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(tx)

	fmt.Println("Sync Products: ", len(products))

	for _, product := range products {

		// Mapping data to elastic formatter
		doc := map[string]interface{}{
			"id":          product.ID,
			"name":        product.Name,
			"principal":   product.Principal,
			"description": product.Description,
			"image":       product.Image,
			"company_id":  product.CompanyID,
			"available":   product.Available,
			"created_at":  product.CreatedAt,
			"updated_at":  product.UpdatedAt,
		}

		req := elastic.NewBulkIndexRequest().Index("products").Id(product.ID).Doc(doc)
		bulkRequest = bulkRequest.Add(req)
	}

	// Execute bulk request
	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		helper.SendErrorResponse(c, http.StatusNotFound, "bulk request failed : %"+err.Error())
	}

	fmt.Printf("Indexed %d documents to Elasticsearch\n", len(bulkResponse.Items))

}
