package handler

import (
	"net/http"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/service"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	srv *service.Manager
}

func NewCategoryHandler(srv *service.Manager) *CategoryHandler {
	return &CategoryHandler{srv: srv}
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	var req models.Category
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ID, err := h.srv.Category.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": ID,
	})
}

func (h *CategoryHandler) ListCategories(c echo.Context) error {
	resp, err := h.srv.Category.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CategoryHandler) GetCategory(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Category.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	id := c.Param("id")

	var req models.Category
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Category.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Category.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
