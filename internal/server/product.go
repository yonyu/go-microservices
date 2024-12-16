package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yonyu/go-microservices/internal/dberrors"
	"github.com/yonyu/go-microservices/internal/models"
)

// what should we do here?

// implement a func to get all products

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	// get all products from the database
	// return them as a slice of Product
	// if there is an error, return it
	// if there is no error, return nil
	vendorId := ctx.QueryParam("vendorId")
	products, err := s.DB.GetAllProducts(ctx.Request().Context(), vendorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)
}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Product)

	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)

		default:
			return ctx.JSON(http.StatusInternalServerError, err)

		}
	}
	return ctx.JSON(http.StatusCreated, product)
}

func (s *EchoServer) GetProductById(ctx echo.Context) error {
	id := ctx.Param("id")
	product, err := s.DB.GetProductById(ctx.Request().Context(), id)

	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, product)
}

func (s *EchoServer) UpdateProduct(ctx echo.Context) error {
	product := new(models.Product)
	id := ctx.Param("id")

	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	if id != product.ProductID {
		return ctx.JSON(http.StatusBadRequest, "id on path doesn't match id on body")
	}

	product, err := s.DB.UpdateProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusInternalServerError, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, product)
}

func (s *EchoServer) DeleteProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	err := s.DB.DeleteProduct(ctx.Request().Context(), id)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.NoContent(http.StatusResetContent)
}
