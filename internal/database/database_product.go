package database

import (
	"context"
	"github.com/yonyu/go-microservices/internal/models"
)

// GetAllProducts(ctx.Request().Context(), vendorId)
func (c Client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error) {
	var products []models.Product

	result := c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorID}).
		Find(&products)
	return products, result.Error
}
