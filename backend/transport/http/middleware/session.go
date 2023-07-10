package middleware

import "github.com/labstack/echo/v4"

type SessionAuth struct {
}

func (s *SessionAuth) SignInMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return next
}

func (s *SessionAuth) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return next
}
