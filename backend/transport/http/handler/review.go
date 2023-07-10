package handler

import (
	"net/http"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/service"
	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	srv *service.Manager
}

func NewReviewHandler(srv *service.Manager) *ReviewHandler {
	return &ReviewHandler{srv: srv}
}

func (h *ReviewHandler) CreateReview(c echo.Context) error {
	productID := c.Param("id")

	var req models.Review
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req.Product.ID = productID

	ID, err := h.srv.Review.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": ID,
	})
}

func (h *ReviewHandler) ListReviews(c echo.Context) error {
	productID := c.Param("id")

	resp, err := h.srv.Review.List(c.Request().Context(), productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ReviewHandler) GetReview(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Review.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ReviewHandler) UpdateReview(c echo.Context) error {
	id := c.Param("id")

	var req models.Review
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Review.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *ReviewHandler) DeleteReview(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Review.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}
