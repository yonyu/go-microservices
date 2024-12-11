package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yonyu/go-microservices/internal/dberrors"
	"github.com/yonyu/go-microservices/internal/models"
)

func (s *EchoServer) GetAllVendors(ctx echo.Context) error {
	vendors, err := s.DB.GetAllVendors(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}

func (s *EchoServer) AddVendor(ctx echo.Context) error {
	vendor := new(models.Vendor)

	if err := ctx.Bind(vendor); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	vendor, err := s.DB.AddVendor(ctx.Request().Context(), vendor)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, vendor)
}
