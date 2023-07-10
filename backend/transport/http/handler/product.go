package handler

import (
	"net/http"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/service"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	srv *service.Manager
}

func NewProductHandler(srv *service.Manager) *ProductHandler {
	return &ProductHandler{srv: srv}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var req models.Product
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ID, err := h.srv.Product.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": ID,
	})
}

func (h *ProductHandler) ListProducts(c echo.Context) error {
	queries := new(models.Queries)
	queries.Name = c.QueryParam("name")
	queries.StartPrice = c.QueryParam("startPrice")
	queries.EndPrice = c.QueryParam("endPrice")
	queries.StartRating = c.QueryParam("startRating")
	queries.EndRating = c.QueryParam("endRating")

	resp, err := h.srv.Product.List(c.Request().Context(), queries)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Product.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	var req models.Product
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Product.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Product.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}
