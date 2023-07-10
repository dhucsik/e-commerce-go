package handler

import (
	"net/http"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/service"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	srv *service.Manager
}

func NewCartHandler(srv *service.Manager) *CartHandler {
	return &CartHandler{srv: srv}
}

func (h *CartHandler) AddToCart(c echo.Context) error {
	var req models.Cart
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ID, err := h.srv.Cart.AddProductToCart(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": ID,
	})
}

func (h *CartHandler) GetUsersCart(c echo.Context) error {
	resp, err := h.srv.Cart.GetUsersCart(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CartHandler) DeleteFromCart(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Cart.DeleteProductFromCart(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
