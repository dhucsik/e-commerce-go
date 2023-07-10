package handler

import (
	"net/http"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv *service.Manager
}

func NewUserHandler(srv *service.Manager) *UserHandler {
	return &UserHandler{srv: srv}
}

func (h *UserHandler) SignUp(c echo.Context) error {
	var req models.User
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ID, err := h.srv.User.SignUp(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": ID,
	})
}

func (h *UserHandler) SignIn(c echo.Context) error {
	var req models.AuthUser
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userData, err := h.srv.User.SignIn(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Set(string(models.ContextKey), userData)

	return nil
}

func (h *UserHandler) UpdatePassword(c echo.Context) error {
	var req models.UpdatePasswordReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.srv.User.UpdatePassword(c.Request().Context(), &req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
