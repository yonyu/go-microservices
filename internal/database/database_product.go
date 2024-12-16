package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/yonyu/go-microservices/internal/dberrors"
	"github.com/yonyu/go-microservices/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (c Client) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).
		Model(&products).
		Clauses(clause.Returning{}).
		Where(models.Product{ProductID: product.ProductID}).
		Updates(models.Product{
			Price:    product.Price,
			Name:     product.Name,
			VendorID: product.VendorID,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "Product", ID: product.ProductID}
	}
	return &products[0], nil
}

func (c Client) DeleteProduct(ctx context.Context, productId string) error {
	result := c.DB.WithContext(ctx).
		Delete(&models.Product{ProductID: productId})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return &dberrors.NotFoundError{Entity: "Product", ID: productId}
	}
	return nil
}
