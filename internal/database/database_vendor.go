package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/yonyu/go-microservices/internal/dberrors"
	"github.com/yonyu/go-microservices/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor

	result := c.DB.WithContext(ctx).
		Find(&vendors)
	return vendors, result.Error
}

func (c Client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	vendor.VendorID = uuid.NewString()

	result := c.DB.WithContext(ctx).Create(&vendor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return vendor, nil
}

func (c Client) GetVendorById(ctx context.Context, vendorId string) (*models.Vendor, error) {
	vendor := &models.Vendor{}

	result := c.DB.WithContext(ctx).
		Where(models.Vendor{VendorID: vendorId}).
		First(&vendor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{}
		}
		return nil, result.Error
	}
	return vendor, nil
}
