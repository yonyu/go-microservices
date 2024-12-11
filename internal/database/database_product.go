package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/yonyu/go-microservices/internal/dberrors"
	"github.com/yonyu/go-microservices/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error) {
	var products []models.Product

	result := c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorID}).
		Find(&products)
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}

	return product, nil
}

func (c Client) GetProductById(ctx context.Context, productId string) (*models.Product, error) {
	product := new(models.Product)

	result := c.DB.WithContext(ctx).
		Where(models.Product{ProductID: productId}).
		First(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{}
		}
		return nil, result.Error
	}

	return product, nil
}
