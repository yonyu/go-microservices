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

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).Find(&services)
	return services, result.Error
}

func (c Client) AddService(ctx context.Context, service *models.Service) (*models.Service, error) {
	service.ServiceId = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&service)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return service, nil
}

func (c Client) GetServiceById(ctx context.Context, serviceId string) (*models.Service, error) {
	service := new(models.Service)

	result := c.DB.WithContext(ctx).Where(models.Service{ServiceId: serviceId}).First(&service)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{}
		}
		return nil, result.Error
	}
	return service, nil
}

func (c Client) UpdateService(ctx context.Context, service *models.Service) (*models.Service, error) {
	var services []models.Service

	result := c.DB.WithContext(ctx).
		Model(&services).
		Clauses(clause.Returning{}).
		Where(models.Service{ServiceId: service.ServiceId}).
		Updates(models.Service{
			Name:  service.Name,
			Price: service.Price,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "Service", ID: service.ServiceId}
	}
	return &services[0], nil
}

func (c Client) DeleteService(ctx context.Context, serviceId string) error {
	result := c.DB.WithContext(ctx).Delete(&models.Service{ServiceId: serviceId})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return &dberrors.NotFoundError{Entity: "Service", ID: serviceId}
	}
	return nil
}
