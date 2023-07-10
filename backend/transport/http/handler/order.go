package handler

import (
	"log"
	"net/http"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/service"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	srv *service.Manager
}

func NewOrderHandler(srv *service.Manager) *OrderHandler {
	return &OrderHandler{srv: srv}
}

func (h *OrderHandler) MakeOrder(c echo.Context) error {
	var req models.Order
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Order.MakeOrder(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	mail := &models.Mail{
		Receiver: req.User.Email,
		Message:  "Ordered successfully",
	}

	log.Println(mail)

	err = h.srv.Mail.SendMail(c.Request().Context(), mail)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *OrderHandler) ListOrders(c echo.Context) error {
	resp, err := h.srv.Order.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *OrderHandler) GetOrder(c echo.Context) error {
	id := c.Param("id")

	resp, err := h.srv.Order.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	id := c.Param("id")

	var req models.Order
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.srv.Order.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	id := c.Param("id")

	err := h.srv.Order.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
