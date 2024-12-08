package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
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
