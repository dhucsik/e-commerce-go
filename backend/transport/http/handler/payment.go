package handler

/*
import (
	"net/http"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/labstack/echo/v4"
)

func (h Manager) CreatePayment(c echo.Context) error {
	var req models.Payment
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ID, err := h.srv.Payment.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": ID,
	})
}

func (h Manager) ListPayments(c echo.Context) error {
	resp, err := h.srv.Payment.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) GetPayment(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Payment.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Manager) UpdatePayment(c echo.Context) error {
	id := c.Param("id")

	var req models.Payment
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Payment.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h Manager) DeletePayment(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Payment.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
*/
